package commands

import (
	"fmt"
	"os/exec"

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
			out, err := exec.Command("docker", "run", "--rm", "hello-world").Output()
			if err != nil {
				return err
			}
			fmt.Fprintln(dockerCli.Out(), string(out))
			return nil
		},
	}
	return cmd
}
