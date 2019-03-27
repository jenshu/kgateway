package secret

import (
	"context"
	"fmt"

	"github.com/solo-io/gloo/projects/gloo/cli/pkg/common"

	envoyutil "github.com/envoyproxy/go-control-plane/pkg/util"
	"github.com/solo-io/gloo/pkg/cliutil"
	"github.com/solo-io/gloo/projects/gloo/cli/pkg/cmd/options"
	"github.com/solo-io/gloo/projects/gloo/cli/pkg/helpers"
	"github.com/solo-io/gloo/projects/gloo/cli/pkg/surveyutils"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	"github.com/solo-io/solo-kit/pkg/errors"
	"github.com/solo-io/solo-projects/projects/gloo/pkg/api/v1/plugins/extauth"

	gloov1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	"github.com/spf13/cobra"
)

func ExtAuthOathCmd(opts *options.Options) *cobra.Command {
	meta := &opts.Metadata
	input := extauth.OauthSecret{}
	cmd := &cobra.Command{
		Use:   "oauth",
		Short: `Create an OAuth secret with the given name (Enterprise)`,
		Long:  `Create an OAuth secret with the given name. The OAuth secrets contains the client_secret as defined in RFC 6749. This is an enterprise-only feature.`,
		RunE: func(c *cobra.Command, args []string) error {
			if len(args) == 1 {
				meta.Name = args[0]
			}
			if opts.Top.Interactive {
				// and gather any missing args that are available through interactive mode
				if err := oauthSecretArgsInteractive(meta, &input); err != nil {
					return err
				}
			}
			// create the secret
			if err := createOauthSecret(opts.Top.Ctx, *meta, input, opts.Create.DryRun); err != nil {
				return err
			}

			return nil
		},
	}

	flags := cmd.Flags()
	flags.StringVar(&input.ClientSecret, "client-secret", "", "oauth client secret")

	return cmd
}

func oauthSecretArgsInteractive(meta *core.Metadata, input *extauth.OauthSecret) error {
	if err := surveyutils.InteractiveNamespace(&meta.Namespace); err != nil {
		return err
	}

	if err := cliutil.GetStringInput("Name of secret:", &meta.Name); err != nil {
		return err
	}

	if err := cliutil.GetStringInput("Enter Client Secret:", &input.ClientSecret); err != nil {
		return err
	}

	return nil
}

func createOauthSecret(ctx context.Context, meta core.Metadata, input extauth.OauthSecret, dryRun bool) error {
	if input.ClientSecret == "" {
		return fmt.Errorf("client-secret not provided")
	}

	secretStruct, err := envoyutil.MessageToStruct(&input)
	if err != nil {
		return errors.Wrapf(err, "Error marshalling oauth secret")
	}

	secret := &gloov1.Secret{
		Metadata: meta,
		Kind: &gloov1.Secret_Extension{
			Extension: &gloov1.Extension{
				Config: secretStruct,
			},
		},
	}

	if dryRun {
		return common.PrintKubeSecret(ctx, secret)
	}

	secretClient := helpers.MustSecretClient()
	if _, err := secretClient.Write(secret, clients.WriteOpts{Ctx: ctx}); err != nil {
		return err
	}
	fmt.Printf("Created OAuth secret [%v] in namespace [%v]\n", meta.Name, meta.Namespace)

	return nil
}
