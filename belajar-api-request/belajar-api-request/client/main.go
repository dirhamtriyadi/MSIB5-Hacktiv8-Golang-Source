package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
)

func getTodos() {
	apiEndpoint := "https://jsonplaceholder.typicode.com/todos"

	req, err := http.NewRequest(http.MethodGet, apiEndpoint, nil)

	if err != nil {
		log.Panicf("error while creating api request instance: %s\n", err.Error())
	}

	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		log.Panicf("error while executing the api request instance: %s\n", err.Error())
	}

	responseBody, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Panicf("error while reading response body from server: %s\n", err.Error())
	}

	fmt.Println("responseBody:", string(responseBody))
}

func createOrder() {
	apiEndpoint := "http://localhost:8080/orders"

	requestBodyJSON := `
		{
			"customerName": "John Doe",
			"items": [
			  {
				"description": "BMW",
				"itemCode": "889",
				"quantity": 13
			  }
			],
			"orderedAt": "2023-07-10T21:21:46+00:00"
		}
	`

	req, err := http.NewRequest(
		http.MethodPost,
		apiEndpoint,
		bytes.NewBuffer([]byte(requestBodyJSON)),
	)

	if err != nil {
		log.Panicf("error while creating api request instance: %s\n", err.Error())
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		log.Panicf("error while executing the create order api: %s\n", err.Error())
	}

	responseBody, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Panicf("error while reading response body: %s\n", err.Error())
	}

	fmt.Println("response from create order api:", string(responseBody))
}

func main() {
	apiEndpoint := "http://localhost:8080"

	body := &bytes.Buffer{}

	writer := multipart.NewWriter(body)

	writer.WriteField("name", "john doe")

	imageFile, err := os.Open("mountain.jpeg")

	if err != nil {
		log.Panicf("error while opening image file: %s\n", err.Error())
	}

	defer imageFile.Close()

	part, err := writer.CreateFormFile("image-file", imageFile.Name())

	_, err = io.Copy(part, imageFile)

	if err != nil {
		log.Panicf("error while copying image file to buffer: %s\n", err.Error())
	}

	req, err := http.NewRequest(http.MethodPost, apiEndpoint, io.NopCloser(body))

	if err != nil {
		log.Panicf("error while creating api request instance: %s\n", err.Error())
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())

	err = writer.Close()

	if err != nil {
		log.Panicf("error while closing the writer instance: %s\n", err.Error())
	}

	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		log.Panicf("error while sending image file via tcp network: %s\n", err.Error())
	}

	_ = resp

}
