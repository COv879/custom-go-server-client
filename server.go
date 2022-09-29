package main

import (
	"fmt"
	"net"
)

const BUFFERSIZE = 1024

func serv() {
	s, err := net.Listen("tcp", "localhost:8180") // s = server
	if err != nil {
		fmt.Println("Error en la conexion:", err)
		return
	}
	defer s.Close()
	fmt.Println("Servidor Iniciado")
	for {
		c, err := s.Accept() // c = client
		if err != nil {
			fmt.Println("Error del cliente", err)
			continue
		}
		go clientHandler(c)
	}
}

func clientHandler(c net.Conn) {
	b := make([]byte, 100) // b = 100 bit array
	bs, err := c.Read(b)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("Mensaje: ", b[:bs])
		fmt.Println("Bystes", bs)
	}

}

func main() {
	go serv()

	var input string
	fmt.Scan(&input)
}
