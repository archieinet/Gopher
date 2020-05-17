package main

import (
	"fmt"
)

type user struct { //defind object, name user, this can be outside, as of now it's inner scope
	ID        int //fields
	FirstName string
	LastName  string
}

func main() {

	var u user
	u.ID = 7
	u.FirstName = "Archie"
	fmt.Println("I'm", u.FirstName)
	u.LastName = "Pliansaneh"

	fmt.Println(u)

	user2 := user{ID: 8, FirstName: "yUi", LastName: "Sakulchot"}
	fmt.Println(user2)

}
