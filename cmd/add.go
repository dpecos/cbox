package cmd

import (
	"bufio"
	"fmt"
	"os"

	"github.com/dpecos/cmdbox/db"
	"github.com/dpecos/cmdbox/models"
	"github.com/spf13/cobra"
)

func readString(label string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(label)
	val, _ := reader.ReadString('\n')
	return val
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new command to your cmdbox",
	Run: func(cmd *cobra.Command, args []string) {

		cmdboxDB := db.Load(dbPath)
		defer cmdboxDB.Close()

		command := models.Cmd{
			Cmd:         readString("Command / snippet: "),
			Title:       readString("Title: "),
			Description: readString("Description: "),
			Url:         readString("URL: "),
		}

		db.Add(command)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
