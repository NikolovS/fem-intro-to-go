package main

import "fmt"

// User is a user type
type User struct {
	ID                         int
	FirstName, LastName, Email string
}
type Group struct {
	role string
	users []User
	newestUser User
	spaceAvailable bool

}

func describeUser(u User) string {
	desc := fmt.Sprintf("Name: %s %s, Email: %s", u.FirstName, u.LastName, u.Email)
	return desc
}

func (user *User) describe() string{
	desc := fmt.Sprintf("Name: %s %s, Email: %s", user.FirstName, user.LastName, user.Email)
	return desc
}


func describeGroup (g Group) string {

	if len(g.users)>2 {
		g.spaceAvailable =false
	}

desc := fmt.Sprintf("This group has %d users. The newest user is %s %s . Accepting new users: %t" , len(g.users),  g.newestUser.FirstName, g.newestUser.LastName, g.spaceAvailable )
return desc
}
func (group *Group) describe() string {
	if len(group.users)>2 {
		group.spaceAvailable =false
	}

desc := fmt.Sprintf("This group has %d users. The newest user is %s %s . Accepting new users: %t" , len(group.users),  group.newestUser.FirstName, group.newestUser.LastName, group.spaceAvailable )
return desc
}

func main() {
	user := User{ID: 1, FirstName: "Marilyn", LastName: "Monroe", Email: "marilyn.monroe@gmail.com"}
	user2 := User{ID: 2, FirstName: "Bo", LastName: "Mo", Email: "bo.mo@gmail.com"}
	
	group := Group{
		role: "admins",
		users: []User{user, user2},
		spaceAvailable: false,
		newestUser: user2,
	}
	// desc := describeUser(user)
	desc := user.describe()
	descGroup :=group.describe()
	
	fmt.Println(desc)
	fmt.Println(descGroup)
	 
}
