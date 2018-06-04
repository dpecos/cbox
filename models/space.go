package models

type Space struct {
	Name    string    `json:"name"`
	Title   string    `json:"title"`
	Entries []Command `json:"entries"`
}

func (space *Space) CommandAdd(command Command) {
	space.Entries = append(space.Entries, command)
}

func (space *Space) CommandList(tag string) []Command {
	if tag == "" {
		return space.Entries
	}

	var result []Command
	for _, command := range space.Entries {
		for _, t := range command.Tags {
			if tag == t {
				result = append(result, command)
			}
		}
	}
	return result
}
