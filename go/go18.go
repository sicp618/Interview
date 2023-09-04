package main

import (
	"fmt"
	"net"
	"net/netip"

	"golang.org/x/exp/constraints"
)

func Sum[V constraints.Integer | constraints.Float](s ...V) (r V) {
	for _, v := range s {
		r += v
	}
	return
}

func runTemplateFunc() {
	fmt.Println("f return ", Sum(1, 2, 3))
	fmt.Println("f return ", Sum(1.0, 2.3, 3.5))
}

func runIpType() {
	ip := "192.168.0.1"
	fmt.Println(net.ParseIP(ip))

	addr, _ := netip.ParseAddr(ip)
	fmt.Println(addr)
}

func main() {
	runTemplateFunc()
	runIpType()
}