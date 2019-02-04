package cli

func optionalSelector(args []string, idx int) *string {
	if len(args) > idx {
		return &args[idx]
	} else {
		return nil
	}
}
