package handlers

import (
	utils "ecommerce/Utils"
	"ecommerce/config"
	"ecommerce/database"
	"encoding/json"
	"log"
	"net/http"
)

func LoginUser(w http.ResponseWriter, r *http.Request) {
	var reqLog database.ReqLogin

	json.NewDecoder(r.Body).Decode(&reqLog)

	usr := database.Login(reqLog)
	if usr == nil {
		log.Println("Login Failed")
		return
	}
	cnf := config.GetConfig()
	accesToken, err := utils.CreateJwt(cnf.SecretKey, utils.Payload{
		FirstName:   usr.FirstName,
		LastName:    usr.LastName,
		Email:       usr.Email,
		IsShopOwner: usr.IsShopOwner,
	})

	if err != nil {
		http.Error(w, "Authentication Failed", 401)
		return
	}
	utils.SendData(w, accesToken, 200)
}
