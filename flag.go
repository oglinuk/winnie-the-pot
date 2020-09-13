package main

import "flag"

var (
	host          = flag.String("h", "", "Host ip")
	port          = flag.Uint("p", 22, "Port")
	hostKey       = flag.String("hk", "", "Path to host rsa key")
	serverVersion = flag.String("sv", "", "Version of server (must start with \"SSH-2.0-\")")
)
