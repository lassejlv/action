package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Prompt(question string) string {
    reader := bufio.NewReader(os.Stdin)
    fmt.Print(question)
    answer, _ := reader.ReadString('\n')
    return strings.TrimSpace(answer)
}
