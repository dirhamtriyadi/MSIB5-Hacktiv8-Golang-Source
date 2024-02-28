package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

var port = ":8080"

type Product struct {
	Id    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type GetAllProductResponse struct {
	Data []Product `json:"data"`
}

type GetOneProductResponse struct {
	Data Product `json:"data"`
}

var products []Product = []Product{}

func handleError(w http.ResponseWriter, errMessage string, statusCode int) {
	w.WriteHeader(statusCode)

	w.Write([]byte(errMessage))

}

func readProduct(w http.ResponseWriter, productId string) {
	id, err := strconv.Atoi(productId)

	if err != nil {
		handleError(w, "invalid product id", http.StatusBadRequest)
		return
	}

	var p *Product

	for _, eachProduct := range products {
		if eachProduct.Id == id {
			p = &eachProduct
			break
		}
	}

	if p == nil {
		handleError(w, "product not found", http.StatusNotFound)
		return
	}

	response := GetOneProductResponse{
		Data: *p,
	}

	b, err := json.Marshal(response)

	if err != nil {
		handleError(w, "something went wrong", http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	w.Write(b)

}

func readProducts(w http.ResponseWriter, r *http.Request) {

	productIdPayload := r.URL.Query().Get("id")

	if productIdPayload != "" {
		readProduct(w, productIdPayload)
		return
	}

	response := GetAllProductResponse{
		Data: products,
	}

	b, err := json.Marshal(response)

	if err != nil {
		w.WriteHeader(500)
		errMessage := "something went wrong"

		w.Write([]byte(errMessage))
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(200)

	w.Write(b)
}

func createProduct(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Content-Type") != "application/json" {
		w.WriteHeader(http.StatusBadRequest)

		errMessage := "invalid request content type"

		w.Write([]byte(errMessage))

		return
	}

	b, err := ioutil.ReadAll(r.Body)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		errMessage := "invalid request payload"

		w.Write([]byte(errMessage))

		return
	}

	var p Product

	err = json.Unmarshal(b, &p)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		errMessage := "invalid request json"

		w.Write([]byte(errMessage))

		return
	}

	if len(products) == 0 {
		p.Id = 1
	} else {
		p.Id = products[len(products)-1].Id + 1
	}

	products = append(products, p)

	responseByte, err := json.Marshal(p)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		errMessage := "something went wrong"

		w.Write([]byte(errMessage))

		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	w.Write(responseByte)
}

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		if r.Method == http.MethodGet {
			readProducts(w, r)
			return
		}

		if r.Method == http.MethodPost {
			createProduct(w, r)
			return
		}

	})

	fmt.Println("Listening on port", port)
	http.ListenAndServe(port, nil)
}
