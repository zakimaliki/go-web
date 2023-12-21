package main

import (
	"fmt"
	"golang-test/src/config"
	"golang-test/src/helper"
	"golang-test/src/routes"
	"net/http"

	"github.com/subosito/gotenv"
)

func main() {
	gotenv.Load()
	config.InitDB()
	helper.Migrate()
	defer config.DB.Close()
	routes.Route()
	fmt.Println("starting web server at http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}
