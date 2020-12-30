package main

import (
	"fmt"
	"net"
	"sort"
)

var startPort int
var endPort int

func worker(hostname string, ports, results chan int) {

	for p := range ports {
		addr := fmt.Sprintf("%s:%d", hostname, p)
		conn, err := net.Dial("tcp", addr)

		if err != nil {
			results <- 0
			continue
		}

		conn.Close()
		results <- p
	}
}

func scanSinglePort(hostname string, port int) {
	addr := fmt.Sprintf("%s:%d", hostname, port)
	conn, err := net.Dial("tcp", addr)

	if err != nil {
		fmt.Printf("closed port: %d\n", port)
		return
	}

	conn.Close()
	fmt.Printf("open port: %d\n", port)
}

func main() {
	PrintBanner()
	hostValid, portsValid := ParseArgs()

	if hostValid == nil && portsValid == nil {
		return
	}

	var hostname string
	var portSlice = portsValid[true]

	if len(portSlice) == 0 {
		return
	}

	startPort = portSlice[0]
	endPort = portSlice[len(portSlice)-1]

	for key, val := range hostValid {
		if !val {
			return
		}

		hostname = key
	}

	if len(portsValid[true]) == 1 {
		scanSinglePort(hostname, startPort)
		return
	}

	ports := make(chan int)
	results := make(chan int)
	workerNum := 100
	var openports []int

	for i := 0; i < workerNum; i++ {
		go worker(hostname, ports, results)
	}

	go func() {
		for i := startPort; i <= endPort; i++ {
			ports <- i
		}
	}()

	for i := startPort; i <= endPort; i++ {
		port := <-results

		if port != 0 {
			openports = append(openports, port)
		}
	}

	close(ports)
	close(results)
	sort.Ints(openports)
	for _, port := range openports {
		fmt.Printf("%d : open\n", port)
	}
}
