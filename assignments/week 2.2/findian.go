package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func checkString(x string) string {
	if strings.Contains(x, "a") && x[0] == 'i' && x[len(x)-3] == 'n' {
		return "Found!" // character at len(x)-1 is new line character and at len(x)-2 is carriage return
		// That is the reason for len(x)-3
	}
	return "Not Found!"
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Input your string: ")
	x, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("Error occuerd: %s", err)
	} else {
		x = strings.ToLower(x)
		fmt.Printf(checkString(x))
	}
}
