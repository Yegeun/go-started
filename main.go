package main

import (
	"fmt"
	"github.com/Yegeun/go-started/models"
)

func main() {
	u := models.User{
		ID:        2
		FirstName: "Alex",
		LastName:  "Hopkin",
	}
	fmt.Println(u)
}
