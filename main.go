package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func getColorCode(color string) string {
	switch color {
	case "black":
		return "\x1b[30m" // Black
	case "red":
		return "\x1b[31m" // Red
	case "green":
		return "\x1b[32m" // Green
	case "yellow":
		return "\x1b[33m" // Yellow
	case "blue":
		return "\x1b[34m" // Blue
	case "magenta":
		return "\x1b[35m" // Magenta
	case "pink":
		return "\x1b[35m" // Pink
	case "cyan":
		return "\x1b[36m" // Cyan
	case "white":
		return "\x1b[37m" // White
	case "brightRed":
		return "\x1b[91m" // Bright Red
	case "brightGreen":
		return "\x1b[92m" // Bright Green
	case "brightYellow":
		return "\x1b[93m" // Bright Yellow
	case "brightBlue":
		return "\x1b[94m" // Bright Blue
	case "brightMagenta":
		return "\x1b[95m" // Bright Magenta
	case "brightCyan":
		return "\x1b[96m" // Bright Cyan
	case "brightWhite":
		return "\x1b[97m" // Bright White
	default:
		return "" // Unknown color, return empty string
	}
}

func contains(arr []int, target int) bool {
	for _, a := range arr {
		if a == target {
			return true
		}
	}
	return false
}

func stringToASCII(s string) []int {
	asciiValues := make([]int, len(s))

	for i, char := range s {
		asciiValues[i] = int(char)
	}

	return asciiValues
}

func splitArray(arr []int, splitNum int) [][]int {
	var result [][]int
	currentSection := []int{}

	for _, num := range arr {
		if num == splitNum {
			result = append(result, currentSection)
			currentSection = []int{}
		} else {
			currentSection = append(currentSection, num)
		}
	}

	// Append the last section
	result = append(result, currentSection)

	return result
}

func main() {
	// declare a var for the file name
	var colorFlag string
	// setting the flag
	flag.StringVar(&colorFlag, "color", "default", "This is the color")

	// parse the flag
	flag.Parse()

	// easy access to args
	args := flag.Args()

	var lettersToBeColored string

	var inputString string

	var typeOfAscii string

	if len(flag.Args()) == 3 {
		// string of characters to be colored from cmd
		lettersToBeColored = args[0]

		// Access the arg
		inputString = args[1]

		// getting the type
		typeOfAscii = args[2]
	} else if len(flag.Args()) == 2 {
		// Access the arg
		inputString = args[0]

		// getting the type
		typeOfAscii = args[1]
	} else {
		fmt.Println("Not the correct arguments")
		fmt.Println("Usage: go run . --color=<color> <letters to be colored> something <Banner>")
		fmt.Println("Usage: go run . --color=<color> something <Banner>")
	}

	// get the color code based on input
	colorCode := getColorCode(colorFlag)

	// turning the string to an array of ascii representation
	arrayOfLettersToBeColoredInASCII := stringToASCII(lettersToBeColored)

	// Replace the escape sequence "\n" with an actual newline character
	inputString = strings.ReplaceAll(inputString, "\\n", "\n")

	// trim and to lower
	typeOfAscii = strings.Trim(typeOfAscii, "")
	typeOfAscii = strings.ToLower(typeOfAscii)

	if typeOfAscii != "shadow" && typeOfAscii != "standard" && typeOfAscii != "thinkertoy" {
		fmt.Println("Please enter a correct format (shadow, standard, thinkertoy)")
		os.Exit(1)
	}

	// trim and to lower
	typeOfAscii = strings.Trim(typeOfAscii, "")
	typeOfAscii = strings.ToLower(typeOfAscii)
	inputString = strings.Trim(inputString, "")

	filename := typeOfAscii + ".txt"

	// Open the file
	file, err := os.Open("./" + filename)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	// turn the string to an array of ascii representation
	arrayOfCharacters := stringToASCII(inputString)

	// split the array based on the 10
	splittedArrayBasedOn10 := splitArray(arrayOfCharacters, 10)

	// Create a new scanner for the file
	scanner := bufio.NewScanner(file)

	// Slice to store lines
	var linesFromFile []string

	// Iterate over each line and add it to the array
	for scanner.Scan() {
		lineFromFile := scanner.Text()
		linesFromFile = append(linesFromFile, lineFromFile)
	}

	// for every array in the array [[65 108 105] [104 101 108 108 111]] 0 & 1
	for j := 0; j < len(splittedArrayBasedOn10); j++ {

		// A 9 time for loop since every character has 8 lines and a new line
		for i := 1; i <= 9; i++ {

			// for every character in the array, (this will repeat 9 times due to outer loop)
			for k, asciiRep := range splittedArrayBasedOn10[j] {

				// read the asciiRep to get the position of the pointer for the lines array
				positionOfpointer := (asciiRep-32)*9 + i

				// check if asciiRep exists in the array of inputted string of chars to be colored
				if contains(arrayOfLettersToBeColoredInASCII, asciiRep) {
					fmt.Print(colorCode + linesFromFile[positionOfpointer] + "\x1b[0m")
				} else if lettersToBeColored == "" {
					fmt.Print(colorCode + linesFromFile[positionOfpointer] + "\x1b[0m")

				} else {
					// print out every line without a new line
					fmt.Print(linesFromFile[positionOfpointer])
				}

				// if we reach the end of an array, print a new line
				if k == len(splittedArrayBasedOn10[j])-1 {
					fmt.Println()
				}
			}
		}
		// if a specific array in the bigger array is empty, output a new line
		if len(splittedArrayBasedOn10[j]) == 0 {
			fmt.Println()
		}

	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

}
