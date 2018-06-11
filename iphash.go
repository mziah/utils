// Author: Masum Z. Hasan

package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func iphash(ip string) int {

	var ipint int

	s := strings.Split(ip, ".")
	//fmt.Printf("spl = %s\n", s)

	for i := range s {
		v, err := strconv.Atoi(s[i])
		if err != nil {
			return 0
		}
		ipint += v
	}
	return ipint
}

func main() {

	ip1 := "10.200.25.5"
	ip2 := "10.127.27.9"
	ip3 := "10.127.27.10"
	ip4 := "10.127.27.11"
	ip5 := "10.300.25.5"
	ip6 := "10.268.25.121"

	fmt.Printf("ip1 = %d\n", iphash(ip1)%10)
	fmt.Printf("ip2 = %d\n", iphash(ip2)%10)
	fmt.Printf("ip3 = %d\n", iphash(ip3)%10)
	fmt.Printf("ip4 = %d\n", iphash(ip4)%10)
	fmt.Printf("ip5 = %d\n", iphash(ip5)%10)
	fmt.Printf("ip6 = %d\n", iphash(ip6)%10)

	fmt.Printf("ip12 = %d\n", (iphash(ip1)+iphash(ip2))%10)
	fmt.Printf("ip34 = %d\n", (iphash(ip3)+iphash(ip4))%10)
	fmt.Printf("ip56 = %d\n", (iphash(ip5)+iphash(ip6))%10)

	fmt.Printf("ip78 = %d\n", (iphash(os.Args[1])+iphash(os.Args[2]))%10)

}

