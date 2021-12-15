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
	router.Route("/").Method("GET").Handler(homeHandler).Insert()
	router.Route("/param/:id").Method("GET").Handler(routeWithParamsHandler).HandlePreflight().Assign(logger).Insert()

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
