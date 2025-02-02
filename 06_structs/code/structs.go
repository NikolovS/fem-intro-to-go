package main

import "fmt"

// User is a user type
type User struct {
	ID                         int
	FirstName, LastName, Email string
}

// // User is a user type
// type User struct {
// ID					       int
// FirstName, LastName, Email string
// }
type Group struct {
	role string
	users []User
	newestUser User
	spaceAvailable bool
}



func describeUser (u User) string {
	desc := fmt.Sprintf("My name is %s %s. My emails is %s", u.FirstName , u.LastName, u.Email)
	return desc
}

func describeGroup (g Group) string {

	if len(g.users)>2 {
		g.spaceAvailable =false
	}

desc := fmt.Sprintf("This group has %d users. The newest user is %s %s . Accepting new users: %t" , len(g.users),  g.newestUser.FirstName, g.newestUser.LastName, g.spaceAvailable )
return desc
}

func main() {
	u := User{ID: 1, FirstName: "Marilyn", LastName: "Monroe", Email: "marilyn.monroe@gmail.com"}
	u2 := User{ID: 2, FirstName: "Blq", LastName: "Blq", Email: "bla.blq@gmail.com"}
	u3 := User{ID: 2, FirstName: "Ha", LastName: "Hahov", Email: "ha.hahov@gmail.com"}

	group := Group{
		role: "admins",
		users: []User{u, u2, u3} ,
		newestUser: u2,
		spaceAvailable: true,
	}
	fmt.Println(describeUser(u))
	fmt.Println(describeGroup(group))
}
