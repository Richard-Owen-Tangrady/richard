package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/Richard-Owen-Tangrady/richard/models"
	"github.com/gorilla/mux"
)

var product []models.Product

func GetProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}

func GetProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	param := mux.Vars(r)
	for _, prod := range product {
		if prod.ProductID == param["productid"] {
			json.NewEncoder(w).Encode(prod)
			return
		}
	}
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var prod models.Product
	_ = json.NewDecoder(r.Body).Decode(&prod)
	product = append(product, prod)
	json.NewEncoder(w).Encode(prod)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	param := mux.Vars(r)
	for index, item := range product {
		if item.ProductID == param["productid"] {
			product = append(product[:index], product[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(product)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	param := mux.Vars(r)
	for index, item := range product {
		if item.ProductID == param["productid"] {
			product = append(product[:index], product[index+1:]...)
			var prod models.Product
			_ = json.NewDecoder(r.Body).Decode(&prod)
			prod.ProductID = param["id"]
			product = append(product, prod)
			json.NewEncoder(w).Encode(prod)
			return
		}
	}
}
