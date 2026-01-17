package main

import (
	"fmt"

	"github.com/FalkSturmfels/dsacalender/pkg/date"
)

func main() {
	fmt.Println("Current Date:", date.CurrentDate.ToString())

	date.SetCurrentDate(69, 3116)

	err := date.SaveCurrentDate(date.CurrentDate)
	if err != nil {
		fmt.Println("Error saving current day:", err)
		return
	}
	fmt.Println("Saved current day successfully.")
	fmt.Println("New Current Date:", date.CurrentDate.ToString())

	fmt.Printf("Day: %+v\n", date.PlusDays(3).ToString())
}
