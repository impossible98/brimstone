package utils

import (
	// import built-in packages
	"fmt"
	"time"
)

func ShowTip(time time.Time, port string) {
	clearScreen()
	fmt.Println()
	fmt.Println(tip(time))
	fmt.Println()
	fmt.Println(localNetwork(port))
	ip, err := getLocalIPv4()
	if err != nil {
	} else {
		fmt.Println(PublicNetwork(ip, port))
	}
	fmt.Println()
}
