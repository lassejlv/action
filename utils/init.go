package utils

import (
	"bufio"
	"fmt"
	"os"
)

func Init() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Welcome to Actionfile!")
	fmt.Println("Enter the name of your first command: ")

	commandName, _ := reader.ReadString('\n')

	fmt.Println("Enter the command string: ")
	commandString, _ := reader.ReadString('\n')

	WriteCommand(commandName, commandString)
	fmt.Println("Command added successfully! Run action --list to see your commands")
}

func WriteCommand(commandName string, commandString string) {
	file, err := os.OpenFile(ConfigFileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = file.WriteString(fmt.Sprintf("%s=%s\n", commandName, commandString))
	if err != nil {
		panic(err)
	}
}
