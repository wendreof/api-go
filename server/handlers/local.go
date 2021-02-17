package handlers

import (
	"fmt"
	"net/http"

	"github.com/wendreof/hello-go/api/models"
)

//Local all HTTP requisitions
func Local(resp http.ResponseWriter, req *http.Request) {
	local := models.Local{}
	// phoneCode, err := strconv.Atoi(req.URL.Path[7:])
	// if err != nil {
	// 	http.Error(resp, "Number not valid", http.StatusBadRequest)
	// 	fmt.Println("[Local] Error on convert number")
	// }

	if err := ModelLocal.ExecuteTemplate(resp, "local.html", local); err != nil {
		http.Error(resp, "Was not possibile proceed", http.StatusInternalServerError)
		fmt.Println("[Basic] Error on execute model")
	}
}
