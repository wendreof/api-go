package main

import (
	"fmt"
	"net/http"

	handler "github.com/wendreof/hello-go/api/server/handlers"
)

func main() {
	http.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(resp, "Hello from Go!")
	})

	http.HandleFunc("/Handler", handler.Handler)
	http.HandleFunc("/Basic", handler.Basic)
	fmt.Println("Server available at localhost:8181")
	http.ListenAndServe(":8181", nil)
}
