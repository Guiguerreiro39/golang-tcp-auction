package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

func main() {
	args := os.Args
	if len(args) == 1 {
		fmt.Println("Please provide port number.")
		return
	}

	PORT := ":" + args[1]
	listener, err := net.Listen("tcp", PORT)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer listener.Close()

	con, err := listener.Accept()
	if err != nil {
		fmt.Println(err)
		return
	}

	reader := bufio.NewReader(con)

	for {
		netData, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println(err)
			return
		}
		if strings.TrimSpace(string(netData)) == "STOP" {
			fmt.Println("Exiting TCP server!")
			return
		}

		fmt.Print("-> ", string(netData))
		timer := time.Now()
		myTime := timer.Format(time.RFC3339) + "\n"
		con.Write([]byte(myTime))
	}
}
