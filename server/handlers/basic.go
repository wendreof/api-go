package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/wendreof/hello-go/api/models"
)

//Basic is the requisition handler from /basic route
func Basic(resp http.ResponseWriter, req *http.Request) {
	page := models.Page{}
	page.Hour = time.Now().Format("15:04:05")
	if err := ModelBasic.ExecuteTemplate(resp, "basic.html", page); err != nil {
		http.Error(resp, "Was not possibile proceed", http.StatusInternalServerError)
		fmt.Println("[Basic] Error on execute model")
	}
}
