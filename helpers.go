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

func GetParam(r *http.Request, key string) string {
	params, _ := r.Context().Value(ParamsKey).(Params)

	for _, param := range params {
		if param.Key == key {
			return param.Value
		}
	}
	return ""
}

func EmptyHandler(w http.ResponseWriter, r *http.Request) {}
