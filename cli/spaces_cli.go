package cli

import (
	"github.com/dpecos/cbox/tools"
	"github.com/spf13/cobra"
)

var spacesCmd = &cobra.Command{
	Use:   "spaces",
	Args:  cobra.ExactArgs(0),
	Short: "Show available spaces in your cbox",
	Long:  tools.Logo,
	Run:   ctrl.SpacesList,
}

var spacesAddCmd = &cobra.Command{
	Use:   "add",
	Args:  cobra.ExactArgs(0),
	Short: "Add a new space to your cbox",
	Long:  tools.Logo,
	Run:   ctrl.SpacesAdd,
}

var spacesEditCmd = &cobra.Command{
	Use:   "edit",
	Args:  cobra.ExactArgs(1),
	Short: "Edit an space from your cbox",
	Long:  tools.Logo,
	Run:   ctrl.SpacesEdit,
}

var spacesDeleteCmd = &cobra.Command{
	Use:   "delete",
	Args:  cobra.ExactArgs(1),
	Short: "Delete a space from your cbox",
	Long:  tools.Logo,
	Run:   ctrl.SpacesDelete,
}

func init() {
	rootCmd.AddCommand(spacesCmd)
	spacesCmd.AddCommand(spacesAddCmd)
	spacesCmd.AddCommand(spacesEditCmd)
	spacesCmd.AddCommand(spacesDeleteCmd)
}
