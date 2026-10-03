package handlers

import (
	utils "ecommerce/Utils"
	"ecommerce/database"
	"encoding/json"
	"net/http"
)

func LoginUser(w http.ResponseWriter, r *http.Request) {
	var reqLog database.ReqLogin

	json.NewDecoder(r.Body).Decode(&reqLog)

	login := database.Login(reqLog)
	// if login != nil {
	// 	http.Error(w, "Login Failed", http.StatusBadRequest)
	// 	return
	// }
	utils.SendData(w, login, 200)
}
