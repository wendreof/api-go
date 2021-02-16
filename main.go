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

	request, err := http.NewRequest("GET", "https://jsonplaceholder.typicode.com/posts/1", bytes.NewBuffer(content2Sent))
	if err != nil {
		fmt.Println("[main] Was not possible send it 2 server: ",
			err.Error())
		return
	}
	request.SetBasicAuth("fizz", "buzz")
	request.Header.Set("content-type", "application/jsonb; charset=utf8")
	response, err := client.Do(request)
	if err != nil {
		fmt.Println("[main] Was not possible reach the server: ",
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
		post := models.Post{}
		err = json.Unmarshal(body, &post)
		if err != nil {
			fmt.Println("[main] Error on unmarshal json: ",
				err.Error())
			return
		}
		fmt.Printf("Post is %+v\r\n", post)
	}
}
