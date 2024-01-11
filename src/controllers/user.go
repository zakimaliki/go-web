package controllers

import (
	"encoding/json"
	"fmt"
	"golang-test/src/helper"
	"golang-test/src/models"
	"net/http"
	"os"

	"golang.org/x/crypto/bcrypt"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var input models.User
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Invalid request body")
		return
	}
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	Password := string(hashedPassword)

	item := models.User{
		Email:    input.Email,
		Password: Password,
	}

	w.WriteHeader(http.StatusCreated)
	res := models.PostUser(&item)
	var result, _ = json.Marshal(res)
	w.Write(result)
	return
}

func Login(w http.ResponseWriter, r *http.Request) {
	var input models.User
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Invalid request body")
		return
	}
	ValidateEmail := models.FindEmail(&input)
	// fmt.Fprintln(w, ValidateEmail)
	if len(ValidateEmail) == 0 {
		fmt.Fprintf(w, "Email not Found")
		return
	}
	var passwordSecond string
	for _, user := range ValidateEmail {
		passwordSecond = user.Password
	}
	if err := bcrypt.CompareHashAndPassword([]byte(passwordSecond), []byte(input.Password)); err != nil {
		fmt.Fprintf(w, "Password not Found")
		return
	}
	jwtKey := os.Getenv("SECRETKEY")
	token, err := helper.GenerateToken(jwtKey, input.Email)
	item := map[string]string{
		"Email": input.Email,
		"Token": token,
	}
	var result, _ = json.Marshal(item)
	w.Write(result)
	return
}
