package main

import (
	"fmt"
	"net"
	scannerapi "github.com/Sniper10754/snipermapper/api"
	"strconv"
	"syscall"
	"time"
)

var title string = `
__       _              _  _       
(_ __  o |_) _  ____  _ |_)|_) _  __
__)| | | |  (/_ | |||(_||  |  (/_ | 
`

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

func scanPortsWithPrompt(scanner scannerapi.PortScanner, start int, finish int) []scannerapi.ScanResult {
	ch := make(chan scannerapi.ScanResult)
	var ports []scannerapi.ScanResult
	
	go func() {
		ports = scannerapi.ScanPorts(scanner, start, finish, ch)
	}()

	for {
		val, ok := (<- ch)

		if ok {
			fmt.Printf("Scanning %d...\r", val.Port)
		} else {
			break
		}
	}

	return ports
}

func main() {
	fmt.Println(title)

	fmt.Print("歷 Host: ")
	var host string
	fmt.Scanln(&host)


	addr, err := net.LookupHost(host)

	if err != nil {
		fmt.Println(" Network error: " + err.Error())
		syscall.Exit(1)
	}

	fmt.Printf("ﯱ  Host IP: %s\n", addr[0])

	portStart := 0
	portFinish := 80
	timeout := 2

	tcpPortScanner := scannerapi.NewScanner(scannerapi.TCP, host)
	//udpPortScanner := scannerapi.NewScanner(scannerapi.UDP, host)
	
	// Scan duration = timeout x ports to scan (x 2 because we scan tcp and udp)
	fmt.Printf("  Scan maximum duration: %d~ seconds \n", timeout * (portFinish * 2))
	
	fmt.Println("  Scanning TCP ports")
	tcpScanResult := scanPortsWithPrompt(tcpPortScanner, portStart, portFinish)
	
	//fmt.Println("  Scanning UDP ports")
	//udpScanResult := nil //scanPortList("udp", host, portsToScan, 1)

	startTime := time.Now()

	duration := int(time.Since(startTime).Seconds())

	fmt.Println("  Scan lasted " + strconv.Itoa(duration) +"s")

	fmt.Println("  Scan finished, listing results")
	fmt.Println()

	listScanReport("  tcp/", tcpScanResult)
	//listScanReport("  udp/", udpScanResult)

}