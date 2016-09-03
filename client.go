
Skip to content
This repository

    Pull requests
    Issues
    Gist

    @revcozmo

1
0

    7

revcozmo/Golang-ChatRoom forked from dyzdyz010/Golang-ChatRoom
Code
Pull requests 0
Wiki
Pulse
Graphs
Settings
Golang-ChatRoom/client/client.go
7ff6906 on May 23, 2013
@dyzdyz010 dyzdyz010 First commit.
64 lines (54 sloc) 1.11 KB
package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

var writeStr, readStr = make([]byte, 1024), make([]byte, 1024)

func main() {
	var (
		host   = "127.0.0.1"
		port   = "8000"
		remote = host + ":" + port
		reader = bufio.NewReader(os.Stdin)
	)

	con, err := net.Dial("tcp", remote)
	defer con.Close()

	if err != nil {
		fmt.Println("Server not found.")
		os.Exit(-1)
	}
	fmt.Println("Connection OK.")
	fmt.Printf("Enter your name: ")
	fmt.Scanf("%s", &writeStr)
	in, err := con.Write([]byte(writeStr))
	if err != nil {
		fmt.Printf("Error when send to server: %d\n", in)
		os.Exit(0)
	}
	fmt.Println("Now begin to talk!")
	go read(con)

	for {
		writeStr, _, _ = reader.ReadLine()
		if string(writeStr) == "quit" {
			fmt.Println("Communication terminated.")
			os.Exit(1)
		}

		in, err := con.Write([]byte(writeStr))
		if err != nil {
			fmt.Printf("Error when send to server: %d\n", in)
			os.Exit(0)
		}

	}
}

func read(conn net.Conn) {
	for {
		length, err := conn.Read(readStr)
		if err != nil {
			fmt.Printf("Error when read from server. Error:%s\n", err)
			os.Exit(0)
		}
		fmt.Println(string(readStr[:length]))
	}
}

    Contact GitHub API Training Shop Blog About 

    © 2016 GitHub, Inc. Terms Privacy Security Status Help 

