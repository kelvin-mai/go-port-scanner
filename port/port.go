package port

import (
	"fmt"
	"net"
	"strconv"
	"sync"
	"time"
)

type PortResult struct {
	Port    int
	State   bool
	Service string
}

type PortRange struct {
	Start, End int
}

type ScanResult struct {
	hostname string
	ip       []net.IP
	results  []PortResult
}

var common = map[int]string{
	7:    "echo",
	20:   "ftp",
	21:   "ftp",
	22:   "ssh",
	23:   "telnet",
	25:   "smtp",
	43:   "whois",
	53:   "dns",
	67:   "dhcp",
	68:   "dhcp",
	80:   "http",
	110:  "pop3",
	123:  "ntp",
	137:  "netbios",
	138:  "netbios",
	139:  "netbios",
	143:  "imap4",
	443:  "https",
	513:  "rlogin",
	540:  "uucp",
	554:  "rtsp",
	587:  "smtp",
	873:  "rsync",
	902:  "vmware",
	989:  "ftps",
	990:  "ftps",
	1194: "openvpn",
	3306: "mysql",
	5000: "unpn",
	8080: "https-proxy",
	8443: "https-alt",
}

func ScanPort(protocol, hostname, service string, port int, resultChannel chan PortResult, wg *sync.WaitGroup) {
	defer wg.Done()
	result := PortResult{Port: port, Service: service}
	address := hostname + ":" + strconv.Itoa(port)
	conn, err := net.DialTimeout(protocol, address, 1*time.Second)
	if err != nil {
		result.State = false
		resultChannel <- result
		return
	}
	defer conn.Close()
	result.State = true
	resultChannel <- result
	return
}

func ScanPorts(hostname string, ports PortRange) (ScanResult, error) {
	var results []PortResult
	var scanned ScanResult
	var wg sync.WaitGroup

	resultChannel := make(chan PortResult, ports.End-ports.Start)

	addr, err := net.LookupIP(hostname)
	if err != nil {
		return scanned, err
	}
	for i := ports.Start; i <= ports.End; i++ {
		if service, ok := common[i]; ok {
			wg.Add(1)
			go ScanPort("tcp", hostname, service, i, resultChannel, &wg)

		}
	}
	wg.Wait()
	close(resultChannel)
	for result := range resultChannel {
		results = append(results, result)
	}

	scanned = ScanResult{
		hostname: hostname,
		ip:       addr,
		results:  results,
	}
	return scanned, nil
}

func DisplayScanResult(result ScanResult) {
	ip := result.ip[len(result.ip)-1]
	fmt.Printf("Open ports for %s (%s)\n", result.hostname, ip.String())
	for _, v := range result.results {
		if v.State {
			fmt.Printf("%d	%s\n", v.Port, v.Service)
		}
	}
}

func GetOpenPorts(hostname string, ports PortRange) {
	scanned, err := ScanPorts(hostname, ports)
	if err != nil {
		fmt.Println(err)
	} else {
		DisplayScanResult(scanned)
	}
}
