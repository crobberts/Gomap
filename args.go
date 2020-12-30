package main

import (
	"flag"
	"regexp"
	"strconv"
	"strings"
)

//ParseArgs parses command line args
func ParseArgs() (hostValid map[string]bool, portValid map[bool][]int) {
	hostname := flag.String("i", "", "The `ip/ips` to scan")
	port := flag.String("p", "1-1024", "The `port/ports` to scan on given host")
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
		return parsePortRange(*port)
	}

	if strings.Contains(*port, ",") {
		return parsePortRange(*port)
	}

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

func parsePortRange(portsArgs string) map[bool][]int {
	parts := strings.Split(portsArgs, "-")
	startPort, sErr := strconv.Atoi(parts[0])
	endPort, eErr := strconv.Atoi(parts[1])
	var ports = make(map[bool][]int)
	k := make([]int, 0, 0)

	if sErr != nil || eErr != nil || endPort <= startPort || startPort < 1 || endPort > 65535 {
		return genericPortError(ports)
	}

	for i := startPort; i <= endPort; i++ {
		k = append(k, i)
	}

	ports[true] = k
	return ports
}

func genericPortError(ports map[bool][]int) map[bool][]int {
	flag.PrintDefaults()
	ports[false] = nil
	return ports
}
