package main

import (
	"fmt"
	"log"
	"net"
	"strconv"

	"golang.org/x/crypto/ssh"
)

func handleConn(srvrCfg *ssh.ServerConfig, conn net.Conn) {
	defer conn.Close()
	_, chans, reqs, err := ssh.NewServerConn(conn, srvrCfg)
	if err != nil {
		log.Printf("conn.go::initConn::ssh.NewServerConn()::ERROR: %s", err.Error())
	}

	go HandleRequests(conn.RemoteAddr(), "global", reqs)

	for newChan := range chans {
		go HandleChannel(conn.RemoteAddr(), newChan)
	}
}

func getAddr() string {
	var addr string
	port := strconv.Itoa(int(*port))

	if *host != "" {
		addr = fmt.Sprintf("%s:%s", *host, port)
	} else {
		addr = fmt.Sprintf("0.0.0.0:%s", port)
	}

	return addr
}

func initConn(w *Winnie) {
	addr := getAddr()
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Printf("conn.go::initConn::net.Listen(\"tcp\", %s)::ERROR: %s", addr, err.Error())
	}
	log.Printf("SSH Server listening at %s ...", addr)
	defer lis.Close()
	for {
		conn, err := lis.Accept()
		if err != nil {
			log.Printf("conn.go::initConn::lis.Accept()::ERROR: %s", err.Error())
		}
		log.Printf("Client connected: %s", conn.RemoteAddr())
		go handleConn(w.srvrCfg, conn)
	}
}
