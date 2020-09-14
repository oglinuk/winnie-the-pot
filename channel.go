package main

import (
	"io"
	"log"
	"net"

	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/terminal"
)

// HandleChannel of incoming SSH channels
func HandleChannel(remoteAddr net.Addr, newChan ssh.NewChannel) {
	channel, chanRequests, err := newChan.Accept()
	if err != nil {
		log.Printf("channel.go::HandleChannel::newChan.Accept()::ERROR: %s", err.Error())
	}
	defer channel.Close()
	go HandleRequests(remoteAddr, newChan.ChannelType(), chanRequests)

	if newChan.ChannelType() == "session" {
		term := terminal.NewTerminal(channel, "$ ")
		for {
			line, err := term.ReadLine()
			if err != nil {
				if err == io.EOF {
					log.Println("Terminal closed ...")
				} else {
					log.Println("Failed to read from terminal: ", err.Error())
				}
				break
			}

			log.Printf("%s | %s | %s", remoteAddr, newChan.ChannelType(), line)
		}
	} else {
		data := make([]byte, 256)
		for {
			length, err := channel.Read(data)
			if err != nil {
				if err == io.EOF {
					log.Println("Channel closed ...")
				} else {
					log.Println("Failed to read from channel: ", err.Error())
				}
				break
			}

			log.Printf("%s | %s | %s", remoteAddr, newChan.ChannelType(), string(data[:length]))
		}
	}
}
