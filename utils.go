package cgr

func deleteEmpty(s []string) []string {
	var removed []string

	for _, str := range s {
		if str != "" {
			removed = append(removed, str)
		}
	}

	return removed
}
