package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/wendreof/hello-go/api/models"
	"github.com/wendreof/hello-go/api/repositories"
)

//Local all HTTP requisitions
func Local(resp http.ResponseWriter, req *http.Request) {
	local := models.Local{}
	phoneCode, err := strconv.Atoi(req.URL.Path[7:])
	if err != nil {
		http.Error(resp, "Number not valid!", http.StatusBadRequest)
		fmt.Println("[Local] Error on convert number: ", err.Error())
		return
	}

	sql := "SELECT * FROM place WHERE PHONECODE = ?"
	rows, err := repositories.Db.Queryx(sql, phoneCode)
	if err != nil {
		http.Error(resp, "Was not possible get the number", http.StatusBadRequest)
		fmt.Println("[Local] Was not possibile execute the query: ", sql, err.Error())
		return
	}

	for rows.Next() {
		err = rows.StructScan(&local)
		if err != nil {
			http.Error(resp, "Was not possible get the number", http.StatusBadRequest)
			fmt.Println("[Local] Was not possibile bind the DB data on struct: ", err.Error())
			return
		}
	}

	if err := ModelLocal.ExecuteTemplate(resp, "local.html", local); err != nil {
		http.Error(resp, "Was not possibile proceed", http.StatusInternalServerError)
		fmt.Println("[Basic] Error on execute model")
	}
}
