package main

import (
	"net/http"

	"github.com/ChristianStefaniw/cgr-v2"
)

func main() {
	router := cgr.NewRouter()

	router.Route("/test/:id").Method("GET").Handler(testWithParamHandler).Insert()
	router.Route("/test").Method("GET").Handler(testHandler).Insert()
	router.Route("/favicon.ico").Method("GET").Handler(cgr.EmptyHandler).Insert()

	router.Run("8080")
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ok"))
}

func testWithParamHandler(w http.ResponseWriter, r *http.Request) {
	id := cgr.GetParam(r, "id")
	w.Write([]byte(id))
}
