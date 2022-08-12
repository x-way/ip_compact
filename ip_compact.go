package main

import (
	"bufio"
	"fmt"
	"log"
	"net/netip"
	"os"
	"strings"

	"go4.org/netipx"
)

func readFile(name string) []netip.Prefix {
	var prefixes []netip.Prefix
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
		prefix, err := netip.ParsePrefix(line)
		if err != nil {
			log.Fatal(err)
		}
		prefixes = append(prefixes, prefix)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return prefixes
}

func mustIPSet(b netipx.IPSetBuilder) *netipx.IPSet {
	s, err := b.IPSet()
	if err != nil {
		log.Fatal(err)
	}
	return s
}

func main() {
	var prefixes []netip.Prefix
	if len(os.Args) > 1 {
		if os.Args[1] == "-h" || os.Args[1] == "--help" {
			fmt.Println("Usage: cat iplist.txt | ip_compact")
			return
		}
		for _, fn := range os.Args[1:] {
			prefixes = append(prefixes, readFile(fn)...)
		}
	} else {
		prefixes = readFile("-")
	}

	var builder netipx.IPSetBuilder
	for _, prefix := range prefixes {
		builder.AddPrefix(prefix)
	}
	for _, prefix := range mustIPSet(builder).Prefixes() {
		fmt.Println(prefix.String())
	}
}
