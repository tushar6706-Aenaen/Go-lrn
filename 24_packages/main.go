package main

import (
	"fmt"
	
	"github.com/tushar6706-Aenaen/Go-lrn/user"
	"github.com/tushar6706-Aenaen/Go-lrn/auth"
)

func main() {
	auth.LoginWithCredentials("aenaen","secret")
	session:= auth.GetSession()
	fmt.Println("session",session)

	user := user.User{
		Email: "user@email.com",
		Name: "mamau",
	}

	fmt.Println(user.Email,user.Name)
}