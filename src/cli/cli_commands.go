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
	Short:   "List the content of a space in your cbox",
	Long:    tools.Logo,
	Run:     func(cmd *cobra.Command, args []string) { ctrl.CommandList(args) },
}

var addCmd = &cobra.Command{
	Use:     "add",
	Aliases: []string{"a"},
	Args:    cobra.MaximumNArgs(1),
	Short:   "Add a new command to your cbox",
	Long:    tools.Logo,
	Run:     func(cmd *cobra.Command, args []string) { ctrl.CommandAdd(args) },
}

var editCmd = &cobra.Command{
	Use:     "edit",
	Aliases: []string{"e", "ed"},
	Args:    cobra.ExactArgs(1),
	Short:   "Edit a command of your cbox",
	Long:    tools.Logo,
	Run:     func(cmd *cobra.Command, args []string) { ctrl.CommandEdit(args) },
}

var deleteCmd = &cobra.Command{
	Use:     "delete",
	Aliases: []string{"del"},
	Args:    cobra.ExactArgs(1),
	Short:   "Delete a command of your cbox",
	Long:    tools.Logo,
	Run:     func(cmd *cobra.Command, args []string) { ctrl.CommandDelete(args) },
}

var viewCmd = &cobra.Command{
	Use:     "view",
	Aliases: []string{"v"},
	Args:    cobra.ExactArgs(1),
	Short:   "View one command",
	Long:    tools.Logo,
	Run:     func(cmd *cobra.Command, args []string) { ctrl.CommandView(args) },
}

var tagCmd = &cobra.Command{
	Use:     "tag",
	Aliases: []string{"t"},
	Args:    cobra.MinimumNArgs(2),
	Short:   "Add tags to a command",
	Long:    tools.Logo,
	Run:     func(cmd *cobra.Command, args []string) { ctrl.TagsAdd(args) },
}

var untagCmd = &cobra.Command{
	Use:   "untag",
	Args:  cobra.MinimumNArgs(2),
	Short: "Removes tags from a command",
	Long:  tools.Logo,
	Run:   func(cmd *cobra.Command, args []string) { ctrl.TagsRemove(args) },
}

func init() {
	rootCmd.AddCommand(commandsCmd)
	commandsCmd.AddCommand(addCmd)
	commandsCmd.AddCommand(editCmd)
	commandsCmd.AddCommand(deleteCmd)
	commandsCmd.AddCommand(viewCmd)
	commandsCmd.AddCommand(tagCmd)
	commandsCmd.AddCommand(untagCmd)

	commandsCmd.Flags().BoolVarP(&controllers.ShowCommandsSourceFlag, "view", "v", false, "Show all details about commands")
	viewCmd.Flags().BoolVarP(&controllers.SourceOnlyFlag, "src", "s", false, "view only code snippet source code")
}
