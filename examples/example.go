package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/ChristianStefaniw/cgr-v2"
)

func main() {
	router := cgr.NewRouter()
	logger := cgr.NewMiddleware(loggerMiddleware)
	cors := cgr.NewMiddleware(corsMiddleware)
	router.Route("/param/test").Method("GET").Handler(homeHandler).Insert()
	router.Route("/param/test/:id").Method("GET").Handler(routeWithParamsHandler).HandlePreflight().Assign(logger, cors).Insert()
	router.Run("8080")
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "home")
}

func routeWithParamsHandler(w http.ResponseWriter, r *http.Request) {
	id := cgr.GetParam(r, "id")
	fmt.Fprint(w, id)
}

func loggerMiddleware(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Logger middleware executing...")
	fmt.Println(time.Now())
}

func corsMiddleware(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*")

	w.Header().Add("Access-Control-Allow-Headers", "*")
}
