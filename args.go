package main

import (
	"fmt"
	"flag"
)

func ParseArgs() (hostValid, portValid bool){
	fmt.Println("Hello from args.go")
	
	hostname := flag.String("i", "", "The `ip/ips` to scan")
	port := flag.String("p", "", "The `port/ports` to scan on given host")
	flag.Parse()

	return validateHost(hostname), validatePort(port)
}

func validateHost(hostname *string) bool {
	//perform checks
	return false
}

func validatePort(port *string) bool {
	//perform checks
	return false
}

