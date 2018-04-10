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

func readStringMulti(label string) string {
	fmt.Printf("%s (empty line to finish): ", label)

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

func readString(label string) string {
	fmt.Printf("%s: ", label)

	reader := bufio.NewReader(os.Stdin)
	val, _ := reader.ReadString('\n')
	return strings.TrimSpace(val)
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new command to your cmdbox",
	Run: func(cmd *cobra.Command, args []string) {

		cmdboxDB := db.Load(dbPath)
		defer cmdboxDB.Close()

		command := models.Cmd{
			Title:       readString("Title"),
			Description: readStringMulti("Description"),
			URL:         readString("URL"),
			Cmd:         readStringMulti("Command / snippet"),
		}

		db.Add(command)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
