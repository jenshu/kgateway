package create

import (
	"github.com/solo-io/gloo/projects/gloo/cli/pkg/cmd/create"
	"github.com/solo-io/gloo/projects/gloo/cli/pkg/cmd/options"
	"github.com/solo-io/go-utils/cliutils"
	"github.com/spf13/cobra"
)

/*
Implementation of new CLI extension format.
RootCmd is imported by the command which intends on extending the base command.
rootCmd is a pointer to the OS gloo create cmd, and MustReplaceCmd replaces it's
create virtual service command with the one located locally called VSCreate.
*/

func RootCmd(opts *options.Options, optionsFunc ...cliutils.OptionsFunc) *cobra.Command {
	rootCmd := create.RootCmd(opts, optionsFunc...)
	// Replace old command with new one, replaces by name of command.
	cliutils.MustReplaceCmd(rootCmd, VSCreate(opts))
	return rootCmd
}
