package main

import (
	"net/http"

	"github.com/ChristianStefaniw/cgr-v2"
)

func main() {
	router := cgr.NewRouter()

	router.Route("/test", "GET", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})

	router.Run("8080")
}
