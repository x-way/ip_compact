package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"

	"github.com/mikioh/ipaddr"
)

func main() {
	f := os.Stdin
	if len(os.Args) > 1 {
		var err error
		f, err = os.Open(os.Args[1])
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
	}
	scanner := bufio.NewScanner(f)
	var prefixesv6 []ipaddr.Prefix
	var prefixesv4 []ipaddr.Prefix
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if !strings.Contains(line, "/") {
			if strings.Contains(line, ":") {
				line = line + "/128"
			} else {
				line = line + "/32"
			}
		}
		_, ipNet, err := net.ParseCIDR(line)
		if err != nil {
			log.Fatal(err)
		}
		if strings.Contains(line, ":") {
			prefixesv6 = append(prefixesv6, *(ipaddr.NewPrefix(ipNet)))
		} else {
			prefixesv4 = append(prefixesv4, *(ipaddr.NewPrefix(ipNet)))
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	for _, prefix := range ipaddr.Aggregate(prefixesv6) {
		fmt.Println(prefix)
	}
	for _, prefix := range ipaddr.Aggregate(prefixesv4) {
		fmt.Println(prefix)
	}
}
