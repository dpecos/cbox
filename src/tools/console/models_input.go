package console

import (
	"fmt"
	"strings"

	"github.com/dplabs/cbox/src/models"
)

func readTags() ([]string, error) {
	tags := ReadString("Tags (separated by space)")
	tagList := strings.Split(tags, " ")
	for _, tag := range tagList {
		if tag != "" {
			if !CheckValidChars(tag) {
				PrintError(MSG_NOT_VALID_CHARS)
				return nil, fmt.Errorf(MSG_NOT_VALID_CHARS)
			}
		}
	}
	return tagList, nil
}

func ReadCommand(space *models.Space) *models.Command {
	command := models.Command{
		Label:       strings.ToLower(ReadString("Label", ONLY_VALID_CHARS, NOT_EMPTY_VALUES)),
		Description: ReadString("Description"),
		URL:         ReadString("URL"),
		Code:        ReadString("Code / Command", MULTILINE, NOT_EMPTY_VALUES),
		Tags:        []string{},
	}

	var tags []string
	for tags == nil {
		var err error
		tags, err = readTags()
		if err != nil {
			tags = nil
		}
	}
	for _, tag := range tags {
		command.TagAdd(tag)
	}

	command.Selector = space.Selector.CloneForItem(command.Label)

	fmt.Println()

	return &command
}

func EditCommand(command *models.Command) {
	command.Label = strings.ToLower(EditString("Label", command.Label, ONLY_VALID_CHARS, NOT_EMPTY_VALUES))
	command.Description = EditString("Description", command.Description)
	command.URL = EditString("URL", command.URL)
	command.Code = EditString("Code / Command", command.Code, MULTILINE, NOT_EMPTY_VALUES)

	fmt.Println()
}

func ReadSpace() *models.Space {
	space := models.Space{
		Label:       strings.ToLower(ReadString("Label", NOT_EMPTY_VALUES, ONLY_VALID_CHARS)),
		Description: ReadString("Description"),
		Entries:     []*models.Command{},
	}
	space.Selector = models.NewSelector(models.TypeNone, "", space.Label, "")

	fmt.Println()

	return &space
}

func EditSpace(space *models.Space) {
	space.Label = strings.ToLower(EditString("Label", space.Label, NOT_EMPTY_VALUES, ONLY_VALID_CHARS))
	space.Description = EditString("Description", space.Description)

	fmt.Println()
}
