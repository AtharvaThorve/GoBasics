package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	userInfo := make(map[string]string)
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your name: ")
	name, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error occured:", err)
	} else {
		name = name[:len(name)-2]
		userInfo["name"] = name
	}
	fmt.Print("Enter your address: ")
	addr, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error occured:", err)
	} else {
		addr = addr[:len(addr)-2]
		userInfo["address"] = addr
	}
	user, _ := json.Marshal(userInfo)
	fmt.Println("Json object of the map:", string(user))
}
