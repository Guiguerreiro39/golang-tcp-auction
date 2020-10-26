package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	args := os.Args
	if len(args) == 1 {
		fmt.Println("Please provide a host:port string")
		return
	}
	CONNECT := args[1]

	server, err := net.ResolveUDPAddr("udp4", CONNECT)
	con, err := net.DialUDP("udp4", nil, server)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("The UDP server is %s\n", con.RemoteAddr().String())
	defer con.Close()

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">> ")
		text, _ := reader.ReadString('\n')
		data := []byte(text + "\n")
		_, err = con.Write(data)

		if strings.TrimSpace(string(data)) == "STOP" {
			fmt.Println("Exiting UDP client!")
			return
		}

		if err != nil {
			fmt.Println(err)
			return
		}

		buffer := make([]byte, 1024)
		n, _, err := con.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println(err)
			return
		}
		
		fmt.Printf("Reply: %s\n", string(buffer[0:n]))
	}
}
      