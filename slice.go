//Write a program which prompts the user to enter integers and stores the integers in a sorted slice.
//The program should be written as a loop. Before entering the loop, the program should create an empty integer slice
//of size (length) 3. During each pass through the loop, the program prompts the user to enter an integer to be added to
//the slice. The program adds the integer to the slice, sorts the slice, and prints the contents of the slice in sorted
//order. The slice must grow in size to accommodate any number of integers which the user decides to enter.
//The program should only quit (exiting the loop) when the user enters the character ‘X’ instead of an integer.

//Submit your source code for the program, “slice.go”.

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {

	fmt.Println("!!!!!!!!!!!!!!!!!!! Welcome to slice.go !!!!!!!!!!!!!!!!!!!")
	fmt.Println("You can type 'X' to leave at any time.")


	integerSlice := make([]int, 0, 3)
	inputScanner := bufio.NewScanner(os.Stdin)

	for {

		fmt.Println("Please enter an integer value.")

		inputScanner.Scan()
		userInput := inputScanner.Text()

		if userInput == "X" {
			os.Exit(0)
		}


		intValue, err := strconv.ParseInt(userInput, 10, 64)

		if err != nil {
			continue
		}

		integerSlice = append(integerSlice, int(intValue))
		sort.Ints(integerSlice)

		fmt.Println("Your sorted integers thus far ", integerSlice)
	}
}
