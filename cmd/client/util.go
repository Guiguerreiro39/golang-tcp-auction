package main

import (
	"bufio"
	"fmt"
)

func receiver() {
	// Listen to the server response
	in := bufio.NewReader(con)

	message, _ := in.ReadString('\n')
	fmt.Print("->: " + message)
}
