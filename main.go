package main

import (
	"fmt"

	"github.com/kelvin-mai/go-port-scanner/port"
)

func main() {
	fmt.Println("Hello world")

	open := port.ScanPort("tcp", "localhost", 1313)
	fmt.Printf("Port Open: %t\n", open)

	results := port.InitialScan("localhost")
	fmt.Println(results)
}
