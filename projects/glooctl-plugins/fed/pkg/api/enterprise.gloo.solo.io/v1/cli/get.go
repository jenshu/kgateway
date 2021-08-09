// Code generated by skv2. DO NOT EDIT.

package enterprise_gloo_cli_client

import (
	"context"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/olekukonko/tablewriter"
	"github.com/solo-io/gloo/pkg/cliutil"
	rpc_edge_v1 "github.com/solo-io/solo-projects/projects/apiserver/pkg/api/rpc.edge.gloo/v1"
	"github.com/solo-io/solo-projects/projects/glooctl-plugins/fed/pkg/cmd/options"
	"github.com/solo-io/solo-projects/projects/glooctl-plugins/fed/pkg/constants"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

func AuthConfig(opts *options.Options) *cobra.Command {
	cmd := &cobra.Command{
		Use:     constants.AUTH_CONFIG_COMMAND.Use,
		Aliases: constants.AUTH_CONFIG_COMMAND.Aliases,
		Short:   "list AuthConfigs across all clusters",
		Long:    "usage: glooctl fed get authconfig [NAME] [--namespace=namespace] [-o FORMAT]",
		RunE: func(cmd *cobra.Command, args []string) error {
			flag.Parse()
			portFwd, err := cliutil.PortForward(opts.Namespace, "deploy/gloo-fed-console", opts.ApiserverPort, opts.ApiserverPort, false)
			if err != nil {
				return err
			}
			defer func() {
				if portFwd.Process != nil {
					portFwd.Process.Kill()
					portFwd.Process.Release()
				}
			}()
			grpcOpts := []grpc.DialOption{
				grpc.WithInsecure(),
				grpc.WithBlock(),
			}
			serverAddr := "localhost:" + opts.ApiserverPort
			ctx, _ := context.WithTimeout(opts.Ctx, 10*time.Second)
			conn, err := grpc.DialContext(ctx, serverAddr, grpcOpts...)
			if err != nil {
				return err
			}
			defer conn.Close()
			client := rpc_edge_v1.NewEnterpriseGlooResourceApiClient(conn)
			authConfigs, err := client.ListAuthConfigs(opts.Ctx, &rpc_edge_v1.ListAuthConfigsRequest{})
			if err != nil {
				return err
			}

			table := tablewriter.NewWriter(os.Stdout)
			table.SetHeader([]string{"CLUSTER", "NAMESPACE", "NAME"})
			for _, v := range authConfigs.GetAuthConfigs() {
				table.Append([]string{v.GetMetadata().GetClusterName(), v.GetMetadata().GetNamespace(), v.GetMetadata().GetName()})
			}
			if table.NumLines() == 0 {
				fmt.Printf("No resources found.\n")
			} else {
				table.Render()
			}
			return nil
		},
	}
	return cmd
}
