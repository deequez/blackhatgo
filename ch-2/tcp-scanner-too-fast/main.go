package main

import (
	"fmt"
	"net"
)

func main() {
	for i := 1; i < 1024; i++ {
		go func(j int) {
			address := fmt.Sprintf("scanme.nmap.org:%d", j)
			conn, err := net.Dial("tcp", address)
			if err != nil {
				// port is closed or filtered
				return
			}
			conn.Close()
			fmt.Printf("%d open\n", j)
		}(i)
	}
}
