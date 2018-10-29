package main

import (
	"fmt"
	"log"

	user "github.com/smockoro/ShopApp/user"
)

func main() {
	conn := "DB connection"
	uc := user.NewUserController(conn)
	user, err := uc.View(1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(user)
}
