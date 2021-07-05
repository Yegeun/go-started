package models

type User struct {
	ID	 		int
	FirstName 	string
	LastName 	string
}

var(
	users []*User // this is a slice pointer using a pointer means that you can edit this variable all over without copying anything
	nextID = 1 // you don't have to specify the type unless you need e.g. int32

)
