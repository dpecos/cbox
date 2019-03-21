package cli

import (
	"github.com/dplabs/cbox/src/controllers"
	"github.com/dplabs/cbox/src/tools"
	"github.com/spf13/cobra"
)

var commandsCmd = &cobra.Command{
	Use:     "commands",
	Aliases: []string{"command", "c", "cmd", "list", "ls", "l"},
	Args:    cobra.MaximumNArgs(1),
	Short:   "List the content of an space",
	Long:    tools.Logo,
	Run:     func(cmd *cobra.Command, args []string) { ctrl.CommandList(optionalSelector(args, 0)) },
}

var addCmd = &cobra.Command{
	Use:     "add",
	Aliases: []string{"a"},
	Args:    cobra.MaximumNArgs(1),
	Short:   "Add a new command into an space",
	Long:    tools.Logo,
	Run:     func(cmd *cobra.Command, args []string) { ctrl.CommandAdd(optionalSelector(args, 0)) },
}

var editCmd = &cobra.Command{
	Use:     "edit",
	Aliases: []string{"e", "ed"},
	Args:    cobra.ExactArgs(1),
	Short:   "Edit a command an existing command",
	Long:    tools.Logo,
	Run:     func(cmd *cobra.Command, args []string) { ctrl.CommandEdit(args[0]) },
}

var deleteCmd = &cobra.Command{
	Use:     "delete",
	Aliases: []string{"del"},
	Args:    cobra.ExactArgs(1),
	Short:   "Delete a command",
	Long:    tools.Logo,
	Run:     func(cmd *cobra.Command, args []string) { ctrl.CommandDelete(args[0]) },
}

var viewCmd = &cobra.Command{
	Use:     "view",
	Aliases: []string{"v"},
	Args:    cobra.ExactArgs(1),
	Short:   "Display all the details for an specific command",
	Long:    tools.Logo,
	Run:     func(cmd *cobra.Command, args []string) { ctrl.CommandView(args[0]) },
}

var tagCmd = &cobra.Command{
	Use:     "tag",
	Aliases: []string{"t"},
	Args:    cobra.MinimumNArgs(2),
	Short:   "Add tags to an existing command",
	Long:    tools.Logo,
	Run:     func(cmd *cobra.Command, args []string) { ctrl.TagsAdd(args[0], args[1:]...) },
}

var untagCmd = &cobra.Command{
	Use:   "untag",
	Args:  cobra.MinimumNArgs(2),
	Short: "Removes tags from a command",
	Long:  tools.Logo,
	Run:   func(cmd *cobra.Command, args []string) { ctrl.TagsRemove(args[0], args[1:]...) },
}

var copyCmd = &cobra.Command{
	Use:   "copy",
	Args:  cobra.MinimumNArgs(1),
	Short: "Copy a command from one local space into another",
	Long:  tools.Logo,
	Run:   func(cmd *cobra.Command, args []string) { ctrl.CommandCopy(args[0], optionalSelector(args, 1)) },
}

func init() {
	rootCmd.AddCommand(commandsCmd)
	commandsCmd.AddCommand(addCmd)
	commandsCmd.AddCommand(editCmd)
	commandsCmd.AddCommand(deleteCmd)
	commandsCmd.AddCommand(viewCmd)
	commandsCmd.AddCommand(tagCmd)
	commandsCmd.AddCommand(untagCmd)
	commandsCmd.AddCommand(copyCmd)

	commandsCmd.Flags().BoolVarP(&controllers.ShowCommandsSourceFlag, "view", "v", false, "Show all details about commands")
	viewCmd.Flags().BoolVar(&controllers.SourceOnlyFlag, "src", false, "view only code snippet source code")
}
