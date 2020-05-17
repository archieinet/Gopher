package models

type User struct {
	ID        int
	FirstName string
	LastName  string
}

var (
	users  []*User //colleciton with pointer
	nextID = 1     //int type nextID
)
