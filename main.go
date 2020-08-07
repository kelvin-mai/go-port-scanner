package main

import (
	"github.com/kelvin-mai/go-port-scanner/port"
)

func main() {
	// called with url
	scanned, ok := port.GetOpenPorts("www.freecodecamp.com", port.PortRange{Start: 75, End: 85})
	if ok {
		port.DisplayScanResult(scanned)
	}

	// called with ip address
	scanned, ok = port.GetOpenPorts("104.26.10.78", port.PortRange{Start: 8079, End: 8090})
	if ok {
		port.DisplayScanResult(scanned)
	}

	// verbose called with ip address and no host name returned -- single open port
	scanned, ok = port.GetOpenPorts("104.26.10.78", port.PortRange{Start: 440, End: 450})
	if ok {
		port.DisplayScanResult(scanned)
	}

	// verbose called with ip adress and valid host name returned -- single open port
	scanned, ok = port.GetOpenPorts("137.74.187.104", port.PortRange{Start: 440, End: 450})
	if ok {
		port.DisplayScanResult(scanned)
	}

	// verbose called with host name -- multiple ports returned
	scanned, ok = port.GetOpenPorts("scanme.nmap.org", port.PortRange{Start: 20, End: 80})
	if ok {
		port.DisplayScanResult(scanned)
	}
}
