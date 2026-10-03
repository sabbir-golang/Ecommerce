package database

import (
	"fmt"
)

type User struct {
	ID          int    `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"Last_name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	IsShopOwner string `json:"is_shop_owner"`
}
type ReqLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

var users []User

func (u User) Store() User {
	fmt.Println(u.ID)
	if u.ID != 0 {
		fmt.Println("User Already exits")
		return u
	}
	u.ID = len(users) + 1
	users = append(users, u)
	return u
}
func Login(req ReqLogin) *User {
	fmt.Println("in login func", len(users))
	for _, usr := range users {

		fmt.Println("Login ", usr.Email, req.Email)
		if usr.Email == req.Email && usr.Password == req.Password {
			fmt.Println("Login success")
			return &usr
		}
	}
	return nil
}
