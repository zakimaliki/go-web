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

// func Products_data(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	if r.Method == "GET" {
// 		var result, err = json.Marshal(models.Products)

// 		if err != nil {
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 			return
// 		}
// 		w.WriteHeader(http.StatusOK)
// 		w.Write(result)
// 		return
// 	} else if r.Method == "POST" {
// 		var product models.Product
// 		err := json.NewDecoder(r.Body).Decode(&product)
// 		if err != nil {
// 			w.WriteHeader(http.StatusBadRequest)
// 			fmt.Fprintf(w, "Invalid request body")
// 			return
// 		}
// 		models.Products = append(models.Products, product)
// 		w.WriteHeader(http.StatusCreated)
// 		var result, _ = json.Marshal(models.Products)
// 		w.Write(result)
// 		return
// 	}
// 	http.Error(w, "", http.StatusBadRequest)
// }

// func Product_data(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

// 	// Mendapatkan nilai parameter "id" dari path URL
// 	idParam := r.URL.Path[len("/product/"):]

// 	// Mengubah nilai parameter "id" menjadi integer
// 	id, err := strconv.Atoi(idParam)
// 	if err != nil {
// 		http.Error(w, "Invalid product ID", http.StatusBadRequest)
// 		return
// 	}

// 	// Mencari indeks produk berdasarkan ID
// 	var foundIndex = -1
// 	for i, p := range models.Products {
// 		if p.Id == id {
// 			foundIndex = i
// 			break
// 		}
// 	}

// 	// Mengembalikan error jika produk tidak ditemukan
// 	if foundIndex == -1 {
// 		http.Error(w, "Product not found", http.StatusNotFound)
// 		return
// 	}

// 	if r.Method == "GET" {
// 		// Mengonversi produk menjadi JSON
// 		result, err := json.Marshal(models.Products[foundIndex])
// 		if err != nil {
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 			return
// 		}

// 		// Mengirimkan data produk sebagai respons
// 		w.WriteHeader(http.StatusOK)
// 		w.Write(result)
// 	} else if r.Method == "PUT" {
// 		// Mendecode body request menjadi produk
// 		var updatedProduct models.Product
// 		err := json.NewDecoder(r.Body).Decode(&updatedProduct)
// 		if err != nil {
// 			http.Error(w, "Invalid request body", http.StatusBadRequest)
// 			return
// 		}

// 		// Memperbarui produk
// 		models.Products[foundIndex] = updatedProduct

// 		// Mengonversi produk yang diperbarui menjadi JSON
// 		result, _ := json.Marshal(updatedProduct)

// 		// Mengirimkan data produk yang diperbarui sebagai respons
// 		w.WriteHeader(http.StatusOK)
// 		w.Write(result)
// 	} else if r.Method == "DELETE" {
// 		// Menghapus produk dari slice
// 		models.Products = append(models.Products[:foundIndex], models.Products[foundIndex+1:]...)

// 		// Mengirimkan respons berhasil hapus
// 		w.WriteHeader(http.StatusOK)
// 		fmt.Fprint(w, "Product deleted successfully")
// 	} else {
// 		// Mengembalikan error jika metode HTTP tidak diizinkan
// 		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
// 	}
// }
