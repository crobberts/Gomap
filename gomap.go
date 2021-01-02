package main

import (
	"fmt"
	"net"
	"sort"
)

var startPort int
var endPort int
var workerNum int = 100
var portList []int

func worker(hostname string, ports, results chan int) {

	for p := range ports {
		addr := fmt.Sprintf("%s:%d", hostname, p)
		conn, err := net.Dial("tcp", addr)

		fmt.Printf("testing %s:%d ..\n", hostname, p)
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
	ports := make(chan int)
	results := make(chan int)
	var openports []int
	list := false
	if hostValid == nil && portsValid == nil {
		return
	}

	var hostname string

	for key, val := range hostValid {
		if !val {
			return
		}

		hostname = key
	}

	if val, ok := portsValid["single"]; ok {
		startPort = val[0]
		scanSinglePort(hostname, startPort)
		return
	} else if val, ok := portsValid["range"]; ok {
		startPort = val[0]
		endPort = val[len(val)-1]
		list = false
	} else if val, ok := portsValid["list"]; ok {
		list = true
		portList = val
	}

	for i := 0; i < workerNum; i++ {
		go worker(hostname, ports, results)
	}

	if list {
		go func() {
			for _, val := range portList {
				ports <- val
			}
		}()

		for i := 0; i < len(portList); i++ {
			port := <-results

			if port != 0 {
				openports = append(openports, port)
			}
		}

	} else {
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
	}

	close(ports)
	close(results)
	sort.Ints(openports)
	for _, port := range openports {
		fmt.Printf("%d : open\n", port)
	}
}
