package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

const PORT = 6379

func main() {
	fmt.Println("Listening on port :", PORT)

	server, err := net.Listen("tcp", fmt.Sprintf(":%d", PORT))
	if err != nil {
		fmt.Println(err)
		return
	}

	conn, err := server.Accept()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	for {
		buf := make([]byte, 1024)

		_, err = conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("failed reading from Client: ", err.Error())
			os.Exit(1)
		}

		conn.Write([]byte("+OK\r\n"))
	}

}
