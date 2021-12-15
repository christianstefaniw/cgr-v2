<div align='center'>
	<img src='https://raw.githubusercontent.com/christianstefaniw/cgr-v2/master/assets/logo.png'>
</div>

<div align='center'>
	<a href="https://goreportcard.com/report/github.com/ChristianStefaniw/cgr-v2">
		<img src="https://goreportcard.com/badge/github.com/ChristianStefaniw/cgr-v2"/>
	</a>
	<a href="https://img.shields.io/tokei/lines/github/christianstefaniw/cgr-v2">
		<img src="https://img.shields.io/tokei/lines/github/christianstefaniw/cgr-v2">
	</a>
	<a href="https://img.shields.io/github/license/christianstefaniw/cgr-v2">
		<img src="https://img.shields.io/github/license/christianstefaniw/cgr-v2">
	</a>
</div>

### Installation
```go get github.com/ChristianStefaniw/cgr-v2```

### Instructions


### Example
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

```
