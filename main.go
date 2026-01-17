package main

import (
	"fmt"

	"github.com/FalkSturmfels/dsacalender/data"
)

func main() {
	for _, month := range data.Months {
		fmt.Printf("%s (%s)\n", month.Month, month.Real)
		for _, day := range month.Days {
			fmt.Printf("   Name: %s\n", day.Name)
		}
	}
}
