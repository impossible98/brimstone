package utils

import (
	// import built-in packages
	"fmt"
	"net"
	"strconv"
	"time"
)

func tip(time time.Time) string {
	app_name := fmt.Sprintf(green(bold("%s")), APP_NAME)
	version := fmt.Sprintf(green("v%s"), VERSION)
	differenceTime := strconv.FormatInt(GetTime().Sub(time).Milliseconds(), 10)

	return "  " + app_name + " " + version + "  " + gray("ready in") + " " + differenceTime + " ms"
}

func localNetwork() string {
	return bold(green("  ➜")) + "  " + bold("Local") + ":   " + blue("http://localhost:") + blue(port) + blue("/")
}

func PublicNetwork(ip net.IP) string {
	return bold(green("  ➜")) + "  " + bold("Network") + ": " + blue("http://") + blue(ip.String()) + blue(":") + blue(port) + blue("/")
}
