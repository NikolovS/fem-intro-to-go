package main

import "fmt"

// func main() {
// 	var name string
// 	var namePointer *string

// 	fmt.Println("Name:", name)
// 	fmt.Println("Name *:", namePointer)
// }

// // ******************************************************

// func main() {
// 	var name string = "Beyonce"
// 	var namePointer *string = &name
// 	var nameValue = *namePointer

// 	fmt.Println("Name:", name)
// 	fmt.Println("Name *:", namePointer)
// 	fmt.Println("Name Value:", nameValue)

// }

// // ******************************************************

// func changeName(n *string) {
// 	*n = strings.ToUpper(*n)
// }

// func main() {
// 	name := "Elvis"
// 	changeName(&name)
// 	fmt.Println(name)
// }

// // ******************************************************

type User struct {
	ID int
	Name string
	Email string
}

func updateEmail (u *User) {
	 
		u.Email = "gosho@abv.bg"
		fmt.Println("in update email: ", u.Email)
	
}


func main() {
	u:= User{ID: 1, Name: "Stefan" , Email: "stefan@abv.bg"}
	 updateEmail(&u)
    fmt.Println(u)
}
