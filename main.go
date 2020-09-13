package main

import (
	"flag"
	"io"
	"log"
	"os"
)

func init() {
	f, err := os.OpenFile("winnie.log", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0744)
	if err != nil {
		log.Fatalf("main.go::init::os.OpenFile()::ERROR: %s", err.Error())
	}
	mr := io.MultiWriter(os.Stdout, f)
	log.SetOutput(mr)
	flag.Parse()
}

func main() {
	w := NewWinnie()
	initConn(w)
}
