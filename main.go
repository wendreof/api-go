package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	client := &http.Client{
		Timeout: time.Second * 30,
	}

	response, err := client.Get("https://www.google.com.br")
	if err != nil {
		fmt.Println("[main] Was not possible reach the google servers: ",
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

	request, err := http.NewRequest("GET", "https://www.google.com.br", nil)
	if err != nil {
		fmt.Println("[main] Was not possible request it: ",
			err.Error())
		return
	}
	request.SetBasicAuth("teste", "teste")
	response, err = client.Do(request)
	if err != nil {
		fmt.Println("[main] Was not possible reach the google servers: ",
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
