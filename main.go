package main

import (
	"fmt"
	"log"
	"net"

	"github.com/cilium/cilium/pkg/ip"
)

func main() {

	// https://en.wikipedia.org/wiki/Reserved_IP_addresses
	privateCIDRs := []string{
		"10.0.0.0/8",
		"172.16.0.0/12",
		"192.168.0.0/16",
		"100.64.0.0/10",
		"127.0.0.0/8",
		"169.254.0.0/16",
		"224.0.0.0/4",
		"233.252.0.0/24",
	}

	_, allIPs, err := net.ParseCIDR("0.0.0.0/0")
	if err != nil {
		log.Fatal(err.Error())
	}

	parsedCIDRs := []*net.IPNet{}
	for _, privateCIDR := range privateCIDRs {
		_, parsedCIDR, err := net.ParseCIDR(privateCIDR)
		if err != nil {
			log.Fatal(err.Error())
		}

		parsedCIDRs = append(parsedCIDRs, parsedCIDR)
	}

	for _, cidr := range ip.RemoveCIDRs([]*net.IPNet{allIPs}, parsedCIDRs) {
		fmt.Println(cidr)
	}
}
