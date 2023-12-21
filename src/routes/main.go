package routes

import (
	"fmt"
	"golang-test/src/controllers"
	"net/http"
)

func Route() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "hallo")
	})
	http.HandleFunc("/products", controllers.Products_data)
	http.HandleFunc("/product/", controllers.Product_data)
}
