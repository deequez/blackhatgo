package main

import (
	"fmt"
	"net"
)

func main() {
	// network - identifies the kind of connection to initiate
	// address - the host to which you wish to connect. For IPv4/TCP
	//           connections, this string will take the form of host:port.
	_, err := net.Dial("tcp", "scanme.nmap.org:80")
	if err == nil {
		fmt.Println("Connection successful")
	}
}
