package utils

import (
	"fmt"
)

func Init() {

	commandName := Prompt("What is the name of the command? ")
	fmt.Println(commandName)
}
