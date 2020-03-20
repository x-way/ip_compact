# ip_compact
[![CircleCI](https://circleci.com/gh/x-way/ip_compact/tree/master.svg?style=svg)](https://circleci.com/gh/x-way/ip_compact/tree/master)
[![Go Report Card](https://goreportcard.com/badge/github.com/x-way/ip_compact)](https://goreportcard.com/report/github.com/x-way/ip_compact)

"Compact" a list of IP prefixes (removing duplicates, merging adjacent ranges).

## Usage

```
# cat test.txt
192.168.4.0/24
192.168.5.0/24
2001:db8::123
192.168.6.0/23
192.168.5.5
2001:db8::123/128

# cat test.txt | ip_compact
2001:db8::123/128
192.168.4.0/22
```
