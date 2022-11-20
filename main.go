package main

import (
	"fmt"
	"net"
	scannerapi "github.com/Sniper10754/snipermapper/api"
	"strconv"
	"syscall"
	"time"
)

func makeRange(min, max int) []int {
    a := make([]int, max-min+1)
    for i := range a {
        a[i] = min + i
    }
    return a
}

func listScanReport(prefix string, results []scannerapi.ScanResult) {
	for i := 0; i < len(results); i++ {
		scan := results[i]

		if scan.State {
			fmt.Printf("%s Port open: %d\n", prefix, scan.Port)
		}
	}
	
}

func main() {
	fmt.Print("歷 Host: ")
	var host string
	fmt.Scanln(&host)


	addr, err := net.LookupHost(host)

	if err != nil {
		fmt.Println(" Network error: " + err.Error())
		syscall.Exit(1)
	}

	fmt.Printf("ﯱ  Host IP: %s\n", addr[0])

	portsToScan := makeRange(0, 80)
	timeout := 2

	// Scan duration = timeout x ports to scan (x 2 because we scan tcp and udp)
	fmt.Printf(" Scan maximum duration: %d~ seconds \n", timeout * (len(portsToScan) * 2))
	
	startTime := time.Now()

	tcpScanResult := scannerapi.ScanPorts("tcp", host, portsToScan, int32(timeout))
	// udpScanResult := scannerapi.ScanPorts("udp", host, portsToScan, int32(timeout))

	duration := int(time.Since(startTime).Seconds())

	fmt.Println(" Scan lasted " + strconv.Itoa(duration) +"s")

	fmt.Println("  Scan finished, listing results")

	listScanReport(" tcp ->", tcpScanResult)
	// listScanReport(" udp ->", udpScanResult)

}