package cli

import (
	"github.com/dpecos/cbox/tools"
	"github.com/spf13/cobra"
)

var (
	viewSnippet bool
	filterTag   string
	sourceOnly  bool
)

var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"l", "ls"},
	Args:    cobra.MaximumNArgs(1),
	Short:   "List the content of a space in your cbox",
	Long:    tools.Logo,
	Run:     ctrl.CommandList,
}

var addCmd = &cobra.Command{
	Use:   "add",
	Args:  cobra.MaximumNArgs(1),
	Short: "Add a new command to your cbox",
	Long:  tools.Logo,
	Run:   ctrl.CommandAdd,
}

var editCmd = &cobra.Command{
	Use:     "edit",
	Aliases: []string{"e"},
	Args:    cobra.ExactArgs(1),
	Short:   "Edit a command from your cbox",
	Long:    tools.Logo,
	Run:     ctrl.CommandEdit,
}

var deleteCmd = &cobra.Command{
	Use:     "delete",
	Aliases: []string{"del"},
	Args:    cobra.ExactArgs(1),
	Short:   "Delete a command from your cbox",
	Long:    tools.Logo,
	Run:     ctrl.CommandDelete,
}

var viewCmd = &cobra.Command{
	Use:     "view",
	Aliases: []string{"v"},
	Args:    cobra.ExactArgs(1),
	Short:   "View one command",
	Long:    tools.Logo,
	Run:     ctrl.CommandView,
}

func init() {
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(editCmd)
	rootCmd.AddCommand(deleteCmd)
	rootCmd.AddCommand(viewCmd)

	listCmd.Flags().BoolVarP(&viewSnippet, "view", "v", false, "Show all details about commands")
	viewCmd.Flags().BoolVarP(&sourceOnly, "src", "s", false, "view only code snippet source code")
}
