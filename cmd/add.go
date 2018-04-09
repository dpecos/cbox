package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/dpecos/cmdbox/db"
	"github.com/dpecos/cmdbox/models"
	"github.com/spf13/cobra"
)

func readString(label string) string {
	fmt.Print(label)

	arr := make([]string, 0)
	scanner := bufio.NewScanner(os.Stdin)
	for {
		scanner.Scan()
		text := scanner.Text()
		if len(text) != 0 {
			arr = append(arr, text)
		} else {
			break
		}
	}

	val := strings.Join(arr, "\n")
	return strings.TrimSpace(val)
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
			URL:         readString("URL: "),
		}

		db.Add(command)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
