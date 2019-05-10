package ly

import (
	"fmt"
	"strconv"
)

type IPAddr [4]byte

func (p IPAddr) String() string {
	result := string("")
	length := len(p) - 1
	for i, v := range p {
		if i != length {
			result += strconv.Itoa(int(v)) + "."
		} else {
			result += strconv.Itoa(int(v))
		}
	}
	return result
}

func PrintSringer() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
