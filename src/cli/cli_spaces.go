package cli

import (
	"github.com/dplabs/cbox/src/tools"
	"github.com/spf13/cobra"
)

var spacesCmd = &cobra.Command{
	Use:     "spaces",
	Aliases: []string{"space", "s"},
	Args:    cobra.ExactArgs(0),
	Short:   "Show a list of the spaces available in your cbox",
	Long:    tools.Logo,
	Run:     func(cmd *cobra.Command, args []string) { ctrl.SpacesList() },
}

var spacesCreateCmd = &cobra.Command{
	Use:   "create",
	Args:  cobra.ExactArgs(0),
	Short: "Create a new pristine space",
	Long:  tools.Logo,
	Run:   func(cmd *cobra.Command, args []string) { ctrl.SpacesCreate() },
}

var spacesEditCmd = &cobra.Command{
	Use:     "edit",
	Aliases: []string{"e", "ed"},
	Args:    cobra.ExactArgs(1),
	Short:   "Edit an existing space",
	Long:    tools.Logo,
	Run:     func(cmd *cobra.Command, args []string) { ctrl.SpacesEdit(args[0]) },
}

var spacesDestroyCmd = &cobra.Command{
	Use:   "destroy",
	Args:  cobra.ExactArgs(1),
	Short: "Delete an space and all its content from your cbox",
	Long:  tools.Logo,
	Run:   func(cmd *cobra.Command, args []string) { ctrl.SpacesDestroy(args[0]) },
}

func init() {
	rootCmd.AddCommand(spacesCmd)
	spacesCmd.AddCommand(spacesCreateCmd)
	spacesCmd.AddCommand(spacesEditCmd)
	spacesCmd.AddCommand(spacesDestroyCmd)
}
