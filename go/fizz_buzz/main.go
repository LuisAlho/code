package main

import "fmt"

func main() {

	/* COLORS

		colorReset := "\033[0m"

	    colorRed := "\033[31m"
	    colorGreen := "\033[32m"
	    colorYellow := "\033[33m"
	    colorBlue := "\033[34m"
	    colorPurple := "\033[35m"
	    colorCyan := "\033[36m"
	    colorWhite := "\033[37m"

	*/

	// Colors to aply to text
	colorReset := "\033[0m" //Reset the color to original one
	colorRed := "\033[31m"  //Red
	colorBlue := "\033[34m" //Blue

	fmt.Println(string(colorRed), "Fizz Buzz Game")

	/*
	** TODO: Scanner for a number
	 */

	var x int
	fmt.Println(string(colorBlue), "Please choose a number: ", string(colorReset))

	x = 10
	for i := 0; i < x; i++ {
		fmt.Println(i)
	}

}
