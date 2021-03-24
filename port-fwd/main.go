package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func panicOn(err error) {
	if err != nil {
		panic(err)
	}
}

func parseArgs() (string, string, string) {
	if len(os.Args) < 4 {
		fmt.Println("")
		fmt.Println("")
		panic("Usage: ./ThisProgram <ListenAddr> <NextHopAddr> <Protocol> \nExample: ./ThisProgram :8080 127.0.0.1:80 tcp")
	}
	prot := os.Args[3]
	if prot != "tcp" && prot != "unix" {
		panic("Protocol should be tcp or unix, not " + prot)
	}
	return os.Args[1], os.Args[2], os.Args[3]
}

func main() {
	listenAddr, nextHopAddr, prot := parseArgs()
	fmt.Printf("FORWARD: %v[%v] => %v[%v]\n", prot, listenAddr, prot, nextHopAddr)

	// UDP should use net.ListenUDP.
	ln, err := net.Listen(prot, listenAddr)
	panicOn(err)

	for {
		conn, err := ln.Accept()
		panicOn(err)
		go handleRequest(conn, prot, nextHopAddr)
	}
}

func handleRequest(conn net.Conn, prot, nextHopAddr string) {
	fmt.Println("Dialing new connection...")
	proxy, err := net.Dial(prot, nextHopAddr)
	panicOn(err)

	go copyIO(conn, proxy)
	go copyIO(proxy, conn)
}

func copyIO(src, dest net.Conn) {
	defer src.Close()
	defer dest.Close()
	io.Copy(src, dest)
}

