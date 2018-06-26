package cli

import (
	"github.com/spf13/cobra"
)

var cloudCmd = &cobra.Command{
	Use:   "cloud",
	Short: "Login & sync your spaces to the cloud",
}

var cloudLoginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login to cbox cloud using your Github account",
	Run:   ctrl.CloudLogin,
}

func init() {
	rootCmd.AddCommand(cloudCmd)
	cloudCmd.AddCommand(cloudLoginCmd)
}
