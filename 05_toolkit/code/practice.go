package main

import (
	"fem-intro-to-go/05_toolkit/code/exercise_5a_solution/utils"
	"fmt"
)

func calculateImportantData() int {
	totalValue := utils.Add(1, 2, 3, 4, 5)
	return totalValue
}

func main() {
	fmt.Println("Packages!")
	total := calculateImportantData()
	fmt.Println(total)
}