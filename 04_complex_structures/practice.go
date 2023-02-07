// package main

// import "fmt"

// func avarage(args ... float32) float32 {
// 	var sum float32;
// 	// var counter int
// 	var length = len(args)
// 	for _, value := range args {
// 		// counter++
// 		sum += value
// 	}
// 	// return sum/float32(counter)
// 	return sum / float32(length)

// }
// func main() {
//  var average = avarage(2,4, 3,5, 4,7)
//  fmt.Printf("The average is: %f", average)
// }

//Part 01
//  func calcAverage (arr[5] float64) float64{
// 	 var sum float64

// 	for _ , value := range arr {
// 		sum += value
// 	}
// 	return sum/float64(len (arr))
//  }
// func main () {
//   scores   := [5]float64{ 23.423,24.53123,25.32,65.332,232.43}
// fmt.Println(calcAverage(scores))

// }

//Part 02

// func doesPetExist (arr map[string] string, petName string) bool {
// //  if arr[petName] !=  "" {
// // 	return true
// //  }else {
// // 	return false
// //  }
// _ ,ok := arr[petName]
// return ok
// }
// func main () {
// //   petNames := map[string] string{
// // 	  "fido" : "dog",
// //  	  "pesho":"cat",
// //       "bunny" : "rabbit",
// // 	  "misho": "mouse",
// //   }

// var petNames map[string] string = make(map[string] string)

// petNames["pesho"] = "dog";
// petNames["gosho"] = "mouse"

// 	fmt.Println(petNames)
// 	fmt.Println(doesPetExist(petNames, "gosho"  ))
// }
//Part 3
// func list (slice [] string,args... string) []string {

// 	for _, value := range args {
// 		slice = append(slice, value)
// 	}
// return slice
// }
// func main() {
// 	var groceries = [5] string {"water" , "juice", "cola" , "beer", "bread"}
// 	var mySlice [] string =  groceries[ 1 : 3 ]
// 	fmt.Println(mySlice)
// 	fmt.Println(list(mySlice, groceries[3] , groceries[2] ))

// }

