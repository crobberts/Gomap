package main

import (
	"fmt"
	"net"
	"sort"
)

func worker(ports, results chan int) {
	for p := range ports {
		addr := fmt.Sprintf("localhost:%d", p)
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
	hostValid,portValid := ParseArgs()

	if !hostValid {
		fmt.Println("Host is not valid")
		return
	}

	if !portValid {
		fmt.Println("Port/Ports not valid")
		return
	}

	ports := make(chan int)
	results := make(chan int)
	worker_num := 100;

	var openports []int

	for i := 0; i < worker_num; i++ {
		go worker(ports, results);
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