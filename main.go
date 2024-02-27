package main

import (
	"fmt"
	"net"
)

const PORT = 6379

func main() {
	fmt.Println("Listening on port ", PORT)

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
		resp := ParseResp(conn)
		_, err := resp.Read()
		if err != nil {
			fmt.Println(err)
			return
		}

		writer := NewWriter(conn)
		writer.Write(Value{typ: "string", str: "FOI!"})

	}

}
