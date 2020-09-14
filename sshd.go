package main

import (
	"crypto/ed25519"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"golang.org/x/crypto/ssh"
)

// Winnie the honey pot
type Winnie struct {
	srvrCfg *ssh.ServerConfig
}

// TODO: Refactor to add cfg file check after flag check
func getHostKey() ssh.Signer {
	var k ssh.Signer

	if *hostKey != "" {
		pkBytes, err := ioutil.ReadFile(*hostKey)
		if err != nil {
			log.Printf("ssh.go::config::ioutil.ReadFile(%s)::ERROR: %s", *hostKey, err.Error())
		}

		k, err = ssh.ParsePrivateKey(pkBytes)
		if err != nil {
			log.Printf("ssh.go::config::ssh.ParsePrivateKey(%v)::ERROR: %s", pkBytes, err.Error())
		}
	}

	// If hostKey is not given, check if $HOME/.ssh/id_rsa exists
	defaultHK := fmt.Sprintf("%s/.ssh/id_rsa", os.Getenv("HOME"))
	if _, err := os.Stat(defaultHK); err == nil {
		pkBytes, err := ioutil.ReadFile(defaultHK)
		if err != nil {
			log.Printf("ssh.go::config::ioutil.ReadFile(%s)::ERROR: %s", defaultHK, err.Error())
		}

		k, err = ssh.ParsePrivateKey(pkBytes)
		if err != nil {
			log.Printf("ssh.go::config::ssh.ParsePrivateKey(%v)::ERROR: %s", pkBytes, err.Error())
		}
	} else {
		_, pkBytes, err := ed25519.GenerateKey(nil)
		if err != nil {
			log.Printf("ssh.go::config::ed25519.GenerateKey(nil)::ERROR: %s", err.Error())
		}

		k, err = ssh.NewSignerFromSigner(pkBytes)
		if err != nil {
			log.Printf("ssh.go::config::ssh.NewSignerFromSigner(%v)::ERROR: %s", pkBytes, err.Error())
		}

		log.Println("No host key given, generated a temporary ed25519 one ...")
	}

	return k
}

// TODO: Refactor to add cfg file check after flag check
func getServerVersion() string {
	var version string

	if *serverVersion != "" {
		version = *serverVersion
	} else {
		version = "SSH-2.0-OpenSSH_6.6.1p1"
	}

	return version
}

func config() *ssh.ServerConfig {
	cfg := &ssh.ServerConfig{
		PasswordCallback: func(conn ssh.ConnMetadata, password []byte) (*ssh.Permissions, error) {
			// Log the attackers info and let them in
			log.Printf("%d | %s | %s | %s | %s", time.Now().Unix(),
				conn.RemoteAddr(),
				string(conn.ClientVersion()),
				conn.User(),
				string(password))
			return nil, nil
		},
		ServerVersion: getServerVersion(),
	}

	cfg.AddHostKey(getHostKey())

	return cfg
}

// NewWinnie constructor
func NewWinnie() *Winnie {
	return &Winnie{
		srvrCfg: config(),
	}
}
