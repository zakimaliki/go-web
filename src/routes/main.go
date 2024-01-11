package routes

import (
	"fmt"
	"golang-test/src/controllers"
	"golang-test/src/middleware"
	"net/http"
)

func Route() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "hallo")
	})
	http.Handle("/products", middleware.JwtMiddleware(http.HandlerFunc(controllers.Products_data)))
	http.HandleFunc("/register", controllers.Register)
	http.HandleFunc("/login", controllers.Login)
	http.HandleFunc("/product/", controllers.Product_data)
}
