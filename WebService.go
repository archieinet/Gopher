package main

import (
	"fmt"

	"github.com/archieinet/Gopher/models"
)

func main() {

	u := models.User{
		ID:        2,
		FirstName: "archariya",
		LastName:  "pliansaneh",
	}

	fmt.Println("I'm", u.FirstName)
}
