package main

import "fmt"

type IPAddr [4]byte

func (i IPAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", i[0], i[1], i[2], i[3])
}

func main() {
	host := map[string]IPAddr{
		"local":     {127, 0, 0, 0},
		"GoogleDns": {8, 8, 8, 8},
	}

	for name, ip := range host {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
