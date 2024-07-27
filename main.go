package main

import (
	"bufio"
	"fmt"
	"os"
)

type CommandObj struct {
	command     string
	parsingFunc func(string) (int, string)
}

var commandMap = map[string]CommandObj{
	"-c": {
		command:     "len of bytes",
		parsingFunc: fileSize,
	},
	"-l": {
		command:     "len of lines",
		parsingFunc: noOfLines,
	},
	"-w": {
		command:     "no of words",
		parsingFunc: noOfWords,
	},
}

func noOfLines(filename string) (int, string) {
	scanner, err := getScannerForFile(filename)
	if err != nil {
		return 0, "Error occured while opening the file"
	}
	scanner.Split(bufio.ScanLines)

	noOfLines := 0
	for scanner.Scan() {
		noOfLines++
		fmt.Printf("Read line: %s\n", scanner.Text())
	}
	return noOfLines, ""
}

func noOfWords(filename string) (int, string) {
	scanner, err := getScannerForFile(filename)
	if err != nil {
		return 0, "Error occured while opening the file"
	}
	scanner.Split(bufio.ScanWords)

	noOfWords := 0
	for scanner.Scan() {
		noOfWords++
		fmt.Printf("Read line: %s\n", scanner.Text())
	}
	return noOfWords, ""
}

func getScannerForFile(filename string) (*bufio.Scanner, error) {
	fileObj, err := os.Open(filename)
	if err != nil {
		return &bufio.Scanner{}, err
	}
	scanner := bufio.NewScanner(fileObj)
	return scanner, nil
}

func fileSize(filename string) (int, string) {
	bytes, err := os.ReadFile(filename)

	if err != nil {
		return 0, "Not able to read the file"
	}
	return len(bytes), ""
}

func main() {

	args := os.Args

	if len(args) > 2 {
		fmt.Println("you need to specify the file path and command flag")
	}

	executeCommand(args)

}

func executeCommand(args []string) {
	commandObj, ok := commandMap[args[1]]

	if !ok {
		fmt.Println("command not found")
		return
	}

	output, err := commandObj.parsingFunc(args[2])

	if err != "" {
		fmt.Println(err)
		return
	}

	fmt.Printf("%d %s\n", output, os.Args[2])
}
