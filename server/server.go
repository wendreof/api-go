package main

import (
	"fmt"
	"net/http"

	"github.com/wendreof/hello-go/api/repositories"
	handler "github.com/wendreof/hello-go/api/server/handlers"
)

func init() {
	fmt.Println("Let's start up the server")
}

func main() {

	err := repositories.OpenConnectionMySQL()
	if err != nil {
		fmt.Println("DAMN IT! Was not possible connect to the database: ", err.Error())
		return
	}

	http.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(resp, "Hello from Go!")
	})

	http.HandleFunc("/Handler", handler.Handler)
	http.HandleFunc("/Basic", handler.Basic)
	http.HandleFunc("/Local/", handler.Local)
	fmt.Println("Server available at localhost:8181")
	http.ListenAndServe(":8181", nil)
}
