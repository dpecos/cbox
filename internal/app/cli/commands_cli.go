package cli

import (
	"github.com/dplabs/cbox/internal/pkg"
	"github.com/spf13/cobra"
)

var (
	viewSnippet bool
	filterTag   string
	sourceOnly  bool
)

var commandsCmd = &cobra.Command{
	Use:     "commands",
	Aliases: []string{"command", "c", "cmd", "list", "ls", "l"},
	Args:    cobra.MaximumNArgs(1),
	Short:   "List the content of a space in your cbox",
	Long:    pkg.Logo,
	Run:     ctrl.CommandList,
}

var addCmd = &cobra.Command{
	Use:   "add",
	Args:  cobra.MaximumNArgs(1),
	Short: "Add a new command to your cbox",
	Long:  pkg.Logo,
	Run:   ctrl.CommandAdd,
}

var editCmd = &cobra.Command{
	Use:     "edit",
	Aliases: []string{"e", "ed"},
	Args:    cobra.ExactArgs(1),
	Short:   "Edit a command from your cbox",
	Long:    pkg.Logo,
	Run:     ctrl.CommandEdit,
}

var deleteCmd = &cobra.Command{
	Use:     "delete",
	Aliases: []string{"del"},
	Args:    cobra.ExactArgs(1),
	Short:   "Delete a command from your cbox",
	Long:    pkg.Logo,
	Run:     ctrl.CommandDelete,
}

var viewCmd = &cobra.Command{
	Use:     "view",
	Aliases: []string{"v"},
	Args:    cobra.ExactArgs(1),
	Short:   "View one command",
	Long:    pkg.Logo,
	Run:     ctrl.CommandView,
}

var tagCmd = &cobra.Command{
	Use:     "tag",
	Aliases: []string{"t"},
	Args:    cobra.MinimumNArgs(2),
	Short:   "Add tags to a command",
	Long:    pkg.Logo,
	Run:     ctrl.TagsAdd,
}

var untagCmd = &cobra.Command{
	Use:   "untag",
	Args:  cobra.MinimumNArgs(2),
	Short: "Removes tags from a command",
	Long:  pkg.Logo,
	Run:   ctrl.TagsRemove,
}

func init() {
	rootCmd.AddCommand(commandsCmd)
	commandsCmd.AddCommand(addCmd)
	commandsCmd.AddCommand(editCmd)
	commandsCmd.AddCommand(deleteCmd)
	commandsCmd.AddCommand(viewCmd)
	commandsCmd.AddCommand(tagCmd)
	commandsCmd.AddCommand(untagCmd)

	commandsCmd.Flags().BoolVarP(&viewSnippet, "view", "v", false, "Show all details about commands")
	viewCmd.Flags().BoolVarP(&sourceOnly, "src", "s", false, "view only code snippet source code")
}
