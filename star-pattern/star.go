package main

import "fmt"

func main() {
	var star string = "*"

	for i := 0; i < 10; i++ {
		for j := 0; j < i; j++ {
			fmt.Print(star)
		}
		fmt.Println(" ")
	}
}
