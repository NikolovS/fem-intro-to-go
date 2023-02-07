package main

import "fmt"

func main(){

var sentance string = "This is a string"
for index, value := range sentance {
	if index % 2 ==0 {
		fmt.Println( string(value))
	}

}
}