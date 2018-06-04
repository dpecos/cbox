package cli

import (
	"github.com/spf13/cobra"
)

var sourceOnly bool

var showCmd = &cobra.Command{
	Use:     "show",
	Aliases: []string{"s"},
	Args:    cobra.ExactArgs(1),
	Short:   "Show one command",
	Run: func(cmd *cobra.Command, args []string) {
		// cmdboxDB := db.Load(dbPath)
		// defer cmdboxDB.Close()

		// cmdID := tools.StringToInt(args[0])
		// command := db.Find(cmdID)
		// tools.PrintCommand(command, true, sourceOnly)
	},
}

func init() {
	rootCmd.AddCommand(showCmd)
	showCmd.Flags().BoolVarP(&sourceOnly, "src", "s", false, "Show only code snippet source code")
}
