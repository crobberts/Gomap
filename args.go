package main

import (
	"flag"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

//ParseArgs parses command line args
func ParseArgs() (hostValid map[string]bool, portValid map[string][]int) {
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

func validatePort(port *string) map[string][]int {
	var ports = make(map[string][]int)

	if strings.Contains(*port, "-") {
		return parsePortRange(*port)
	}

	if strings.Contains(*port, ",") {
		return parsePortList(*port)
	}

	gPort, err := strconv.Atoi(*port)

	if err != nil {
		return genericPortError()
	}

	if gPort < 1 || gPort > 65535 {
		return genericPortError()
	}

	var p = make([]int, 0)
	l := append(p, gPort)

	ports["single"] = l
	fmt.Println(ports["single"])
	return ports
}

func parsePortRange(portsArgs string) map[string][]int {
	parts := strings.Split(portsArgs, "-")
	startPort, sErr := strconv.Atoi(parts[0])
	endPort, eErr := strconv.Atoi(parts[1])
	var ports = make(map[string][]int)
	k := make([]int, 0, 0)

	if sErr != nil || eErr != nil || endPort <= startPort || startPort < 1 || endPort > 65535 {
		return genericPortError()
	}

	for i := startPort; i <= endPort; i++ {
		k = append(k, i)
	}

	ports["range"] = k
	return ports
}

func parsePortList(portsArgs string) map[string][]int {
	ports := strings.Split(portsArgs, ",")
	p := make(map[string][]int)
	pr := make([]int, 0, 0)

	for i := 0; i < len(ports); i++ {
		conv, err := strconv.Atoi(ports[i])

		if err != nil {
			return genericPortError()
		}

		pr = append(pr, conv)
	}

	p["list"] = pr
	return p
}

func genericPortError() map[string][]int {
	flag.PrintDefaults()
	ports := make(map[string][]int)
	ports["false"] = nil
	return ports
}
