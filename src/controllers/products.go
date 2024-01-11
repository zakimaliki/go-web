package controllers

import (
	"encoding/json"
	"fmt"
	"golang-test/src/models"
	"net/http"
)

func Products_data(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "GET" {
		res := models.SelectAll()
		var result, err = json.Marshal(res.Value)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(result)
		return
	} else if r.Method == "POST" {
		var input models.Product
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Invalid request body")
			return
		}

		item := models.Product{
			Name:  input.Name,
			Price: input.Price,
			Stock: input.Stock,
		}

		w.WriteHeader(http.StatusCreated)
		res := models.Post(&item)
		var result, _ = json.Marshal(res)
		w.Write(result)
		return

	} else {
		http.Error(w, "", http.StatusBadRequest)

	}
}

func Product_data(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := r.URL.Path[len("/product/"):]

	// id, err := strconv.Atoi(idParam)
	// if err != nil {
	// 	http.Error(w, "Invalid product ID", http.StatusBadRequest)
	// 	return
	// }

	if r.Method == "GET" {
		res := models.Select(id)
		result, err := json.Marshal(res.Value)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(result)
	} else if r.Method == "PUT" {
		var input models.Product
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Invalid request body")
			return
		}
		newProduct := models.Product{
			Name:  input.Name,
			Price: input.Price,
			Stock: input.Stock,
		}
		res := models.Updates(id, &newProduct)
		result, _ := json.Marshal(res)
		w.WriteHeader(http.StatusOK)
		w.Write(result)
	} else if r.Method == "DELETE" {
		res := models.Deletes(id)
		result, _ := json.Marshal(res)
		w.WriteHeader(http.StatusOK)
		w.Write(result)
	} else {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
	}
}
