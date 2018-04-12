package tools

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ReadStringMulti(label string) string {
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

func ReadString(label string) string {
	fmt.Printf("%s: ", label)

	reader := bufio.NewReader(os.Stdin)
	val, _ := reader.ReadString('\n')
	return strings.TrimSpace(val)
}
