package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

const (
	HOST = "localhost"
	PORT = "8080"
	TYPE = "tcp"
)

func main() {
	listner, err := net.Listen(TYPE, HOST+":"+PORT)
	if err != nil {
		log.Fatal("Error starting server ", err)
	}
	defer listner.Close()
	log.Println("Listening on " + HOST + ":" + PORT)
	for {
		conn, err := listner.Accept()
		if err != nil {
			log.Fatal("Error accepting: ", err.Error())
		}
		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	message, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}
	fmt.Print("Message recieved from the client: ", string(message))
	conn.Write([]byte(message + "\n"))
	conn.Close()
}
