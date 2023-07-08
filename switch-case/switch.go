package main

import "fmt"

func main() {
	var input string
	for {
		fmt.Println("Enter a character : ")
		fmt.Scanln(&input)
		switch input {
			case "a":
				fmt.Println("Good choice !")
			case "b":
				fmt.Println("Are you sure ? ")
			case "c":
				fmt.Println("Think again. ")
			case "d":
				fmt.Println("You should choose again")
			default:
				fmt.Println("Invalid Input")
		}
		if input == "q" {
			fmt.Println("Alright see you again!")
			break
		}
	}
}
