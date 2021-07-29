package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// testing how to build an object in go using struct

var commitment1 string
var commitment2 string
var commitment3 string

var commitmentsPath = "commitments.txt"

func createFile(){
	//Check if File exists
	var _, err = os.Stat(commitmentsPath)
	if os.IsNotExist(err) {
		var file, err = os.Create(commitmentsPath)
		fmt.Println("File created: ", commitmentsPath)
		writeFile()
		if isError(err) {
			return
		}
		defer file.Close()
	} else {
		readFile()
		saveOrReset()
	}
}


func writeFile() {
	//Open file using RDWR permissions
	var file, err = os.OpenFile(commitmentsPath, os.O_RDWR, 0644)
	if isError(err){
		return
	}
	defer file.Close()

	//write file line by line
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Type commitment 1: ")
	scanner.Scan()
	commitment1 = scanner.Text()
	fmt.Print("Type commitment 2: ")
	scanner.Scan()
	commitment2 = scanner.Text()
	fmt.Print("Type commitment 3: ")
	scanner.Scan()
	commitment3 = scanner.Text()
	_, err = file.WriteString(
		"Commitment 1: " + commitment1 +" \n" + "Commitment 2: " + commitment2 + " \n" + "Commitment 3: " + commitment3 + " \n")
	if isError(err) {
		return
	}

	err = file.Sync()
	if isError(err) {
		return
	}
	fmt.Println("File updated")
}

func readFile(){
	var file, err = os.OpenFile(commitmentsPath, os.O_RDWR, 0644)
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

func deleteFile() {
	var err = os.Remove(commitmentsPath)
	if isError(err) {
		fmt.Println("error")
		return
	}
	fmt.Println("File Deleted")
}


func isError(err error) bool {
	if err != nil {
		fmt.Println()
	}
	return err != nil
}


func saveOrReset(){
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println(	"Would you like to save or reset these commitments?")
	scanner.Scan()
	switch scanner.Text(){
	case "save", "Save":
		fmt.Println("Do your best today!")
	case "reset", "Reset":
		deleteFile()
		fmt.Println()
		createFile()
	default:
		fmt.Println("Incorrect input: Please type Save or Reset.")
		saveOrReset()
	}

}