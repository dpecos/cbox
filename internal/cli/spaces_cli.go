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

var spacesCreateCmd = &cobra.Command{
	Use:   "create",
	Args:  cobra.ExactArgs(0),
	Short: "Creates a new space to your cbox",
	Long:  tools.Logo,
	Run:   ctrl.SpacesCreate,
}

var spacesEditCmd = &cobra.Command{
	Use:   "edit",
	Args:  cobra.ExactArgs(1),
	Short: "Edit an space from your cbox",
	Long:  tools.Logo,
	Run:   ctrl.SpacesEdit,
}

var spacesDestroyCmd = &cobra.Command{
	Use:   "destroy",
	Args:  cobra.ExactArgs(1),
	Short: "Destroys a space from your cbox",
	Long:  tools.Logo,
	Run:   ctrl.SpacesDestroy,
}

func init() {
	rootCmd.AddCommand(spacesCmd)
	spacesCmd.AddCommand(spacesCreateCmd)
	spacesCmd.AddCommand(spacesEditCmd)
	spacesCmd.AddCommand(spacesDestroyCmd)
}
