package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/wendreof/hello-go/api/models"
)

func main() {
	client := &http.Client{
		Timeout: time.Second * 30,
	}

	user := models.User{}
	user.ID = 1
	user.Name = "@wendreof"

	content2Sent, err := json.Marshal(user)
	if err != nil {
		fmt.Println("[main] Error on convert user 2 json: ",
			err.Error())
		return
	}

	// response, err := client.Get("https://www.google.com.br")
	// if err != nil {
	// 	fmt.Println("[main] Was not possible reach the google servers: ",
	// 		err.Error())
	// 	return
	// }

	// defer response.Body.Close()

	// if response.StatusCode == 200 {
	// 	body, err := ioutil.ReadAll(response.Body)
	// 	if err != nil {
	// 		fmt.Println("[main] Was not possible read the content body: ",
	// 			err.Error())
	// 		return
	// 	}
	// 	fmt.Println(string(body))
	// }

	request, err := http.NewRequest("POST", "https://enjyqmzivj2fp.x.pipedream.net/", bytes.NewBuffer(content2Sent))
	if err != nil {
		fmt.Println("[main] Was not possible send it 2 request bin: ",
			err.Error())
		return
	}
	request.SetBasicAuth("fizz", "buzz")
	request.Header.Set("content-type", "application/jsonb; charset=utf8")
	response, err := client.Do(request)
	if err != nil {
		fmt.Println("[main] Was not possible reach the request bin servers: ",
			err.Error())
		return
	}

	defer response.Body.Close()

	if response.StatusCode == 200 {
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println("[main] Was not possible read the content body: ",
				err.Error())
			return
		}
		fmt.Println(string(body))
	}
}
