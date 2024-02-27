package main

import (
	"fmt"
	"net"
	"strings"
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
		value, err := resp.Read()
		if err != nil {
			fmt.Println(err)
			return
		}

		if value.typ != "array" {
			fmt.Println("Expected Array")
			continue
		}

		if len(value.array) == 0 {
			fmt.Println("array length to short")
			continue
		}

		command := strings.ToUpper(value.array[0].bulk)

		writer := NewWriter(conn)

		handler, ok := Handlers[command]

		if !ok {
			fmt.Println("invalid command: ", command)
			writer.Write(Value{typ: "string", str: ""})
			continue
		}

		args := value.array[1:]
		result := handler(args)

		writer.Write(result)

	}

}
