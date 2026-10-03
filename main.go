package main

import (
	utils "ecommerce/Utils"
	"fmt"
)

// func helloHandler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintln(w, "Hello HandleFunc")
// }

//	func AboutHandler(w http.ResponseWriter, r *http.Request) {
//		fmt.Fprintf(w, "I am a software engineer")
//	}

func main() {
	// cmd.Serve()
	// secret := []byte("My-set")
	// message := []byte("Hello-world")

	// h := hmac.New(sha256.New, secret)
	// h.Write(message)
	// text := h.Sum(nil)

	// fmt.Println(text)
	pay := utils.Payload{
		Sub:         45,
		FirstName:   "Sabbir",
		LastName:    "Ahmed",
		Email:       "sabbir@gmail.com",
		IsShopOwner: true,
	}
	jwt, err := utils.CreateJwt("-secret", pay)
	if err != nil {
		panic(err)
	}
	fmt.Println(jwt)
}
