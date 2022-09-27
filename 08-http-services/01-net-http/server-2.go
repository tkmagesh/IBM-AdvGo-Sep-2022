package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Product struct {
	Id   int     `json:"id"`
	Name string  `json:"name"`
	Cost float32 `json:"cost"`
}

var products []Product

func init() {
	products = []Product{
		{101, "Pen", 10},
		{102, "Pencil", 5},
		{103, "Marker", 50},
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hi from go web server! [using DefaultServerMux]"))
}

func productsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		encoder := json.NewEncoder(w)
		if err := encoder.Encode(products); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	case "POST":
		decoder := json.NewDecoder(r.Body)
		var newProduct Product
		if err := decoder.Decode(&newProduct); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		newProduct.Id = len(products) + 101
		products = append(products, newProduct)
		encoder := json.NewEncoder(w)
		w.WriteHeader(http.StatusCreated)
		if err := encoder.Encode(newProduct); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	case "PUT":
		fmt.Println(r.URL)
		w.Write([]byte("The given product is updated"))
	}

}

/*
PUT  - http://localhost:8080/products/100
*/

func main() {
	//http.DefaultServeMux.HandleFunc("/", indexHandler)
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/products", productsHandler)
	http.ListenAndServe(":8080", nil)
}
