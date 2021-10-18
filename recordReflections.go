package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

var dailyReflection string
var listOfReflections = "listOfReflections.txt"

func recordDailyReflection() {
	fmt.Println("What did you learn or accomplish today?")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	dailyReflection = scanner.Text()

	fmt.Println("Today you learned: " + dailyReflection)

	fmt.Println("Save this reflection to your list?")
	saveReflection()

	fmt.Println("Would you like to read your full list of reflections?")
	readListOfReflections()
}

func readListOfReflections() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	switch scanner.Text() {
	case "Yes", "yes":
		var file, err = os.OpenFile(listOfReflections, os.O_RDWR, 0644)
		if isError(err) {
			return
		}
		defer file.Close()

		var text = make([]byte, 1024)
		for {
			_, err = file.Read(text)

			if err == io.EOF {
				break
			}

			if err != nil && err != io.EOF {
				isError(err)
				break
			}
		}
		fmt.Println("List of reflections:")
		fmt.Println(string(text))
	case "No", "no":
		break
	default:
		fmt.Println("Invalid entry. Please type 'Yes' or 'No'")
		readListOfReflections()
	}
}

func saveReflection() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	switch scanner.Text() {
	case "Yes", "yes":
		var file, err = os.OpenFile(listOfReflections, os.O_RDWR, 0644)
		if isError(err) {
			return
		}
		defer file.Close()
		_, err = file.WriteString(dailyReflection + "\n")
		if isError(err) {
			return
		}
		err = file.Sync()
		if isError(err) {
			return
		}
		fmt.Println("List of reflections Updated.")

	case "No", "no":
		fmt.Println("Reflection not added to list")

	default:
		fmt.Println("Invalid entry. Please enter 'Yes' or 'No'")
		saveReflection()
	}
}
