package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

func init() {
	f, err := os.OpenFile(fmt.Sprintf("%d-winnie.log", time.Now().Unix()), os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0744)
	if err != nil {
		log.Printf("main.go::init::os.OpenFile()::ERROR: %s", err.Error())
	}
	mr := io.MultiWriter(os.Stdout, f)
	log.SetOutput(mr)

	flag.Parse()
}

func main() {
	w := NewWinnie()
	initConn(w)
}
