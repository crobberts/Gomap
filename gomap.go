package main

import (
	"fmt"
	"net"
	"sort"
)

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

func main() {
	PrintBanner()
	hostValid, portsValid := ParseArgs()

	var hostname string
	var portSlice = make([]int, 1, 1)

	for key, val := range hostValid {
		if !val {
			return
		}

		hostname = key
	}

	for _, val := range portsValid {
		for i := 0; i < len(portSlice); i++ {
			portSlice[i] = val[i]
			fmt.Println(portSlice)
		}
	}

	ports := make(chan int)
	results := make(chan int)
	workerNum := 100

	var openports []int

	for i := 0; i < workerNum; i++ {
		go worker(hostname, ports, results)
	}

	go func() {
		for i := 1; i <= 1024; i++ {
			ports <- i
		}
	}()

	for i := 0; i < 1024; i++ {
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
