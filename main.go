package main

import (
	"encoding/json"
	"log"
	"net/http"
	"nujsua/thirgolang/model"

	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Test(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Hello, World!"}`))
}

func main() {
	dsn := "root:mysql@tcp(127.0.0.1:3306)/test_demo1?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	r := mux.NewRouter()

	r.HandleFunc("/api/testing", Test).Methods("GET")

	r.HandleFunc("/api/users", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		var users []model.User
		db.Find(&users)
		// return json response
		json.NewEncoder(w).Encode(users)
	}).Methods("GET")

	r.HandleFunc("/api/users/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		var user model.User
		params := mux.Vars(r)
		db.First(&user, params["id"])
		// return json response
		json.NewEncoder(w).Encode(user)
	}).Methods("GET")

	r.HandleFunc("/api/usersrr", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		var user model.User
		_ = json.NewDecoder(r.Body).Decode(&user)
		db.Create(&user)
		// return json response
		json.NewEncoder(w).Encode(user)
	}).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", r))
}
