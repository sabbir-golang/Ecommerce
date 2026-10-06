package handlers

import (
	utils "ecommerce/Utils"
	"ecommerce/database"
	"encoding/json"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var newUser database.User
	json.NewDecoder(r.Body).Decode(&newUser)

	createUser := newUser.Store()
	utils.SendData(w, createUser, 200)
}
