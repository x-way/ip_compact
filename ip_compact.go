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

func readFile(name string) ([]ipaddr.Prefix, []ipaddr.Prefix) {
	var prefixesv6 []ipaddr.Prefix
	var prefixesv4 []ipaddr.Prefix
	var f *os.File
	if name == "-" {
		f = os.Stdin
	} else {
		var err error
		f, err = os.Open(name)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
	}
	scanner := bufio.NewScanner(f)
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
	return prefixesv6, prefixesv4
}

func main() {
	var prefixesv6 []ipaddr.Prefix
	var prefixesv4 []ipaddr.Prefix
	if len(os.Args) > 1 {
		for _, fn := range os.Args[1:] {
			pfx6, pfx4 := readFile(fn)
			prefixesv6 = append(prefixesv6, pfx6...)
			prefixesv4 = append(prefixesv4, pfx4...)
		}
	} else {
		prefixesv6, prefixesv4 = readFile("-")
	}

	for _, prefix := range ipaddr.Aggregate(prefixesv6) {
		fmt.Println(prefix)
	}
	for _, prefix := range ipaddr.Aggregate(prefixesv4) {
		fmt.Println(prefix)
	}
}
