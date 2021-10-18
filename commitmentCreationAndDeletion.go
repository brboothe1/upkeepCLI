package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// Setup commitment file variables for text and file itself

var commitment string
var commitmentFile = "commitment.txt"

// Search for commitmentFile and if not found create new file.

func createCommitmentFile() {
	//Check if File exists
	var _, err = os.Stat(commitmentFile)
	if os.IsNotExist(err) {
		var file, err = os.Create(commitmentFile)
		fmt.Println("File created: ", commitmentFile)
		writeCommitmentFile()
		if isError(err) {
			return
		}
		defer file.Close()
	} else {
		readCommitmentFile()
		saveOrResetCommitmentFile()
	}
}

func writeCommitmentFile() {
	//Open file using RDWR permissions
	var file, err = os.OpenFile(commitmentFile, os.O_RDWR, 0644)
	if isError(err) {
		return
	}
	defer file.Close()

	//write file line by line
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Type your commitment for the day: ")
	scanner.Scan()
	commitment = scanner.Text()
	_, err = file.WriteString(
		"Today's commitment: " + commitment + " \n")
	if isError(err) {
		return
	}

	err = file.Sync()
	if isError(err) {
		return
	}
	fmt.Println("File updated")
}

// Open and read commitmentFile
func readCommitmentFile() {
	var file, err = os.OpenFile(commitmentFile, os.O_RDWR, 0644)
	if isError(err) {
		return
	}
	defer file.Close()

	//Read file line by line
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
	fmt.Println("Reading from file.")
	fmt.Println(string(text))
}

func deleteCommitmentFile() {
	var err = os.Remove(commitmentFile)
	if isError(err) {
		fmt.Println("error")
		return
	}
	fmt.Println("File Deleted")
}

func saveOrResetCommitmentFile() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Would you like to save or reset this commitment?")
	scanner.Scan()
	switch scanner.Text() {
	case "save", "Save":
		fmt.Println("Do your best today!")
	case "reset", "Reset":
		deleteCommitmentFile()
		fmt.Println()
		createCommitmentFile()
	default:
		fmt.Println("Incorrect input: Please type Save or Reset.")
		saveOrResetCommitmentFile()
	}
}

func isError(err error) bool {
	if err != nil {
		fmt.Println()
	}
	return err != nil
}
