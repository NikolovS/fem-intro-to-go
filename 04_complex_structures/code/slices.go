package main

import "fmt"

func main() {

// 	var myArray [5]int
// 	var mySlice []int

// 	myArray[0] = 1
// 	mySlice[0] = 1

// 	fmt.Println(myArray)
// 	fmt.Println(mySlice)

// 	// ***************************

// 	// var myArray [5]int
// 	// var mySlice []int = make([]int, 5)

// 	// fmt.Println(myArray)
// 	// fmt.Println(mySlice)

// 	// ***************************

// 	// var myArray [5]int
// 	// // var mySlice []int = make([]int, 5)
// 	// var mySlice []int = make([]int, 5, 10)
// 	// // var mySlice = make([]int, 5, 10)

// 	// myArray[0] = 1
// 	// mySlice[0] = 1

// 	// fmt.Println(myArray)
// 	// fmt.Println(mySlice)
// 	// fmt.Println(len(mySlice))
// 	// fmt.Println(cap(mySlice))

// 	// ***************************

	// fruitArray := [5]string{"banana", "pear", "apple", "kumquat", "peach"}

	// var splicedFruit []string = fruitArray[1:3] // ==> ["pear", "apple",]

	// totalFruid := append(splicedFruit, "sol", "pepper", "mqy","bau", "sol" , "test" , "12ds","[[dsa")
	
	// fmt.Println(fruitArray)
	// fmt.Println(splicedFruit)
	// fmt.Println(len(splicedFruit))
	// fmt.Println(cap(splicedFruit))
	// fmt.Println("------------------------------")
	// fmt.Println(totalFruid)
	// fmt.Println(len(totalFruid))
	// fmt.Println(cap(totalFruid))
    var myArray [5]int
    // var mySlice []int = make([]int, 5)
    var mySlice []int = make([]int, 5, 10)
    // var mySlice = make([]int, 5)

    myArray[0] = 1
    mySlice[0] = 2

    fmt.Println(myArray)
    fmt.Println(mySlice)
    fmt.Println(len(mySlice))
    fmt.Println(cap(mySlice))
	fmt.Println(len(myArray))
	fmt.Println(cap(myArray))



// 	// ***************************

// 	// SEE SLIDE

// 	// ***************************

// 	// slice1 := []int{1, 2, 3}
// 	// slice2 := append(slice1, 4, 5)

// 	// fmt.Println(slice1, slice2)
// 	// fmt.Println(len(slice1), cap(slice1))
// 	// fmt.Println(len(slice2), cap(slice2))

// 	// ***************************

// 	// originalSlice := []int{1, 2, 3}
// 	// destination := make([]int, len(originalSlice))

// 	// fmt.Println("Before Copy:", originalSlice, destination)

// 	// mysteryValue := copy(destination, originalSlice)

// 	// // fmt.Println("After Copy:", originalSlice, destination, mysteryValue)
 }
