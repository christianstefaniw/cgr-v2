# cgr-v2
A flexible and easy to use Golang router. Take full control of each route in a clean and readable way! Version 2! 

## Installation
```go get github.com/ChristianStefaniw/cgr-v2```

## Example
```golang
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

func corsMiddleware(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*")

	w.Header().Add("Access-Control-Allow-Headers", "*")
}

```
