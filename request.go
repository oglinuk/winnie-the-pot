package main

import (
	"log"
	"net"

	"golang.org/x/crypto/ssh"
)

type exec struct {
	Cmd string
}

type pty struct {
	Term                    string
	Width, Height           uint32
	PixelWidth, PixelHeight uint32
	Modes                   []byte
}

// HandleRequests for all SSH requests
func HandleRequests(remoteAddr net.Addr, channel string, requests <-chan *ssh.Request) {
	for request := range requests {
		var payload interface{} = request.Payload
		var err error

		switch request.Type {
		case "exec":
			execPayload := exec{}
			err = ssh.Unmarshal(request.Payload, &execPayload)
			if err != nil {
				log.Printf("request.go::Handle::exec::ssh.Unmarshal()::ERROR: %s", err.Error())
			}
			payload = execPayload
		case "pty-req":
			ptyPayload := pty{}
			err = ssh.Unmarshal(request.Payload, &ptyPayload)
			if err != nil {
				log.Printf("request.go::Handle::pty-req::ssh.Unmarshal()::ERROR: %s", err.Error())
			}
			payload = ptyPayload
		}
		log.Printf("%s | %s | %s | %s", remoteAddr, channel, request.Type, payload)
		if request.WantReply {
			err := request.Reply(true, nil)
			if err != nil {
				log.Printf("request.go::Handle::request.Reply()::ERROR: %s", err.Error())
			}
		}
	}
}
