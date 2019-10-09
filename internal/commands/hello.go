package commands

import (
	"fmt"
	"os"

	"github.com/docker/cli/cli"
	"github.com/docker/cli/cli/command"
	"github.com/spf13/cobra"
)

func helloCmd(dockerCli command.Cli) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "hello",
		Short: "Print hello world!",
		Args:  cli.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Fprintln(os.Stdout, "Hello World!")
			return nil
		},
	}
	return cmd
}
