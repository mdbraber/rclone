// Package ls provides the ls command.
package ls

import (
	"context"
	"os"
	"fmt"

	"github.com/rclone/rclone/cmd"
	"github.com/rclone/rclone/cmd/ls/lshelp"
	"github.com/rclone/rclone/fs/operations"
	"github.com/spf13/cobra"
)

func init() {
	cmd.Root.AddCommand(commandDefinition)
}

var commandDefinition = &cobra.Command{
	Use:   "ls remote:path",
	Short: `List the objects in the path with size and path.`,
	Long: `
Lists the objects in the source path to standard output in a human
readable format with size and path. Recurses by default.

Eg

    $ rclone ls swift:bucket
        60295 bevajer5jef
        90613 canole
        94467 diwogej7
        37600 fubuwic

` + lshelp.Help,
	Annotations: map[string]string{
		"groups": "Filter,Listing",
	},
	Run: func(command *cobra.Command, args []string) {
		fmt.Println("Running ls A")
		cmd.CheckArgs(1, 1, command, args)
		fmt.Println("Running ls B")
		fsrc := cmd.NewFsSrc(args)
		fmt.Println("Running ls C")
		cmd.Run(false, false, command, func() error {
			fmt.Println("Running ls operations")
			return operations.List(context.Background(), fsrc, os.Stdout)
		})
	},
}
