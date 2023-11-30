package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// type student struct {
// 	ID    string
// 	Name  string
// 	Grade int
// }

type product struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
	Stock int    `json:"stock"`
}

var products = []product{
	product{1, "baju", 200000, 12},
	product{2, "kemeja", 100000, 8},
	product{3, "jeans", 150000, 6},
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "hallo")
	})
	http.HandleFunc("/products", products_data)
	http.HandleFunc("/product/", product_data)
	fmt.Println("starting web server at http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}

func products_data(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "GET" {
		var result, err = json.Marshal(products)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(result)
		return
	} else if r.Method == "POST" {
		var product product
		err := json.NewDecoder(r.Body).Decode(&product)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Invalid request body")
			return
		}
		products = append(products, product)
		w.WriteHeader(http.StatusCreated)
		var result, _ = json.Marshal(products)
		w.Write(result)
		return
	}
	http.Error(w, "", http.StatusBadRequest)
}

func product_data(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Mendapatkan nilai parameter "id" dari path URL
	idParam := r.URL.Path[len("/product/"):]

	// Mengubah nilai parameter "id" menjadi integer
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	// Mencari indeks produk berdasarkan ID
	var foundIndex = -1
	for i, p := range products {
		if p.Id == id {
			foundIndex = i
			break
		}
	}

	// Mengembalikan error jika produk tidak ditemukan
	if foundIndex == -1 {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	if r.Method == "GET" {
		// Mengonversi produk menjadi JSON
		result, err := json.Marshal(products[foundIndex])
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Mengirimkan data produk sebagai respons
		w.WriteHeader(http.StatusOK)
		w.Write(result)
	} else if r.Method == "PUT" {
		// Mendecode body request menjadi produk
		var updatedProduct product
		err := json.NewDecoder(r.Body).Decode(&updatedProduct)
		if err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		// Memperbarui produk
		products[foundIndex] = updatedProduct

		// Mengonversi produk yang diperbarui menjadi JSON
		result, _ := json.Marshal(updatedProduct)

		// Mengirimkan data produk yang diperbarui sebagai respons
		w.WriteHeader(http.StatusOK)
		w.Write(result)
	} else if r.Method == "DELETE" {
		// Menghapus produk dari slice
		products = append(products[:foundIndex], products[foundIndex+1:]...)

		// Mengirimkan respons berhasil hapus
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "Product deleted successfully")
	} else {
		// Mengembalikan error jika metode HTTP tidak diizinkan
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
	}
}
