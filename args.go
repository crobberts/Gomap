package main

import (
	"flag"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

//ParseArgs parses command line args
func ParseArgs() (hostValid map[string]bool, portValid map[bool][]int) {
	hostname := flag.String("i", "", "The `ip/ips` to scan")
	port := flag.String("p", "", "The `port/ports` to scan on given host")
	flag.Parse()

	return validateHost(hostname), validatePort(port)
}

func validateHost(hostname *string) map[string]bool {
	b := []byte(*hostname)
	m, _ := regexp.Match(`^((25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)$`, b)
	var hostnameValid = make(map[string]bool)

	if !m {
		flag.PrintDefaults()
		hostnameValid[*hostname] = false
		return hostnameValid
	}

	hostnameValid[*hostname] = true
	return hostnameValid
}

func validatePort(port *string) map[bool][]int {
	var ports = make(map[bool][]int)

	if strings.Contains(*port, "-") {
		return parsePortRange(*port, ports)
	}

	trimmed := strings.ReplaceAll(*port, " ", "")
	gPorts := strings.Split(trimmed, ",")
	fmt.Println(gPorts)
	gPort, err := strconv.Atoi(*port)

	if err != nil {
		return genericPortError(ports)
	}

	if gPort < 1 || gPort > 65535 {
		return genericPortError(ports)
	}

	var p = make([]int, 0)
	l := append(p, gPort)

	ports[true] = l
	return ports
}

func parsePortRange(portsArgs string, ports map[bool][]int) map[bool][]int {
	parts := strings.Split(portsArgs, "-")
	startPort, sErr := strconv.Atoi(parts[0])
	endPort, eErr := strconv.Atoi(parts[1])

	if sErr != nil || eErr != nil {
		return genericPortError(ports)
	}

	for i := startPort; i <= endPort; i++ {
		fmt.Println(i)

	}
	return ports
}

func genericPortError(ports map[bool][]int) map[bool][]int {
	flag.PrintDefaults()
	ports[false] = nil
	return ports
}
