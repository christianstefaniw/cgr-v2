package cgr

import "net/http"

func deleteEmpty(s []string) []string {
	var removed []string

	for _, str := range s {
		if str != "" {
			removed = append(removed, str)
		}
	}

	return removed
}

func EmptyHandler(w http.ResponseWriter, r *http.Request) {}
