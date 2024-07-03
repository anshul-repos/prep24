package main

import (
	"fmt"
	"net"
	"strings"
)

//validateIPAddress Function:
//Uses net.ParseIP to check if the IP string is a valid IP address.
//If net.ParseIP returns nil, it is neither a valid IPv4 nor a valid IPv6 address.
//Uses strings.Count to differentiate between IPv4 and IPv6 addresses by checking the presence of . or :.

//validateIPv4 Function:
//Splits the string by . and checks if there are exactly four parts.
//Each part is checked to ensure it is a number between 0 and 255 without leading zeros.

//validateIPv6 Function:
//Splits the string by : and checks if there are exactly eight parts.
//Each part is checked to ensure it contains valid hexadecimal digits (0-9, a-f, A-F) and is no longer than four characters.

func validateIP4(IP string) {
	parts := strings.Split(IP, ".")

	// base length check for parts
	if len(parts) != 4 {
		fmt.Printf("%s is not an valid IP4 address", IP)
		return
	}

	// checks for IP segment/part
	for _, part := range parts {

		// check if length or part is not 0 or greater than 3
		if len(part) == '0' || len(part) > 3 {
			fmt.Printf("%s is not an valid IP4 address", IP)
			return
		}

		// Leading Os are not allowed
		if part[0] == '0' && len(part) > 1 {
			fmt.Printf("%s is not an valid IP4 address", IP)
			return
		}

		num := 0

		for i := 0; i < len(part); i++ {
			if part[i] < '0' || part[i] > '9' {
				fmt.Printf("%s is not an valid IP4 address", IP)
				return
			}
			num = num*10 + int(part[i]-'0')
		}

		if num < 0 || num > 255 {
			fmt.Printf("%s is not an valid IP4 address", IP)
			return
		}

		fmt.Printf("%s is valid IP4 address", IP)
		return
	}

}

func validateIP6(IP string) {
	parts := strings.Split(IP, ".")

	// base length check for parts
	if len(parts) != 8 {
		fmt.Printf("%s is not an valid IP6 address", IP)
		return
	}

	// checks for IP segment/part
	for _, part := range parts {

		if len(part) == 0 || len(part) > 4 {
			fmt.Printf("%s is not an valid IP6 address", IP)
			return
		}

		for i := 0; i < len(part); i++ {

			if !(part[i] >= '0' && part[i] < '9' || part[i] >= 'a' && part[i] < 'f' || part[i] >= 'A' && part[i] < 'F') {
				fmt.Printf("%s is not an valid IP6 address", IP)
				return
			}
		}
	}

	fmt.Printf("%s is valid IP6 address", IP)
	return

}

func validateIP(IP string) {

	// check if string is valid IP address
	if net.ParseIP(IP) == nil {
		fmt.Printf("%s is Not a valid IP address", IP)
		return
	}

	// check for IP4 or IP6
	if strings.Count(IP, ":") >= 2 {
		validateIP6(IP)
	} else {
		validateIP4(IP)
	}

}

func main() {

	println()
	validateIP("192.168.1.1")
	println()
	println()
	validateIP("0.0.0.0")
	println()
	println()
	validateIP("2001:0db8:85a3:0000:0000:8a2e:0370:7334")
	println()
	println()

}
