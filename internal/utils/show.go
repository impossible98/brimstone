package utils

import (
	// import built-in packages
	"fmt"
	"time"
)

var port = GetPort()

func ShowTip(time time.Time) {
	clearScreen()
	fmt.Println()
	fmt.Println(tip(time))
	fmt.Println()
	fmt.Println(localNetwork())
	ip, err := getLocalIPv4()
	if err != nil {
	} else {
		fmt.Println(PublicNetwork(ip))
	}
	fmt.Println()
}
