package main

import (
	"flag"
	"fmt"
)

func main() {
	enableDebug := flag.Bool("debug", false, "Enable debug logging")

	flag.Parse()

	if *enableDebug {
		fmt.Print("Debug: ON\n")
	} else {
		fmt.Print("Debug: OFF\n")
	}
}
