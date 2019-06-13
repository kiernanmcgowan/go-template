package main

import (
	"fmt"
	"net/http"

	"github.com/kiernanmcgowan/go-template/routes"
)

func main() {
	fmt.Println("Server process started")
	routes.BindRoutes()
	http.ListenAndServe(":3000", nil)
}
