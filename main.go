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
}

func createScannerObj(filename string) (int, string) {
	fileObj, err := os.Open(filename)
	if err != nil {
		return 0, "Not able to read the file"
	}
	scanner := bufio.NewScanner(fileObj)
	scanner.Split(bufio.ScanLines)
	noOfLines := 1
	for scanner.Scan() {
		noOfLines++
	}
	return noOfLines, ""
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

	if len(args) != 3 {
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
