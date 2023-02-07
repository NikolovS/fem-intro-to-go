package main

import "fmt"
func main(){
	var name, hometown string
	var time int
	var weather bool
	fmt.Println("Type your name here")
	fmt.Scan(&name)
	fmt.Println("Type your hometown here")
	fmt.Scan(&hometown)
	fmt.Println("Type the time lived in your hometown here")
	fmt.Scan(&time)
	fmt.Println("Type is the weather nice here")
	fmt.Scan(&weather)
	fmt.Printf("My name is %s. I am from %s. I have been living there for more than %d years. They say the weather is amazing, which is %t ", name, hometown, time, weather)
}
// import "fmt"

// func main() {
// 	var name, hometown string
// 	var time int
// 	var weather bool

// 	fmt.Println("Type your name")
// 	fmt.Scan(&name)
// 	fmt.Println("Type your hometown")
// 	fmt.Scan(&hometown)
// 	fmt.Println("Type the time lived in your hometown")
// 	fmt.Scan(&time)
// 	fmt.Println("Is the weather is nice (true/false)?")
// 	fmt.Scan(&weather)
// 	fmt.Printf("Hi! My name is %s. I have lived in %s for %d years. They say the weather is amazing, which is %t", name, hometown, time, weather)

// }
