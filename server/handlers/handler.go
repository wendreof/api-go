package handlers

import (
	"fmt"
	"net/http"
)

//Handler all HTTP requisitions
func Handler(resp http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(resp, "This is the requisition handler using package")
}
