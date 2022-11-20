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

func scanPortList(netprotocol scannerapi.NetworkProtocol, host string, ports []int, timeout int) []scannerapi.ScanResult {
	scannerch := make(chan scannerapi.ScanResult, len(ports))
	var scanResult []scannerapi.ScanResult
	
	go func() {
		scanResult = scannerapi.ScanPorts("tcp", host, ports, int32(timeout), &scannerch)
	}()

	for {
		scanresult := <- scannerch

		fmt.Printf("ﴚ  %d -> %s\r", scanresult.Port, strconv.FormatBool(scanresult.State))
		
		if scanresult.Port == ports[len(ports)-1] {
			break
		}
	}

	return scanResult
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

	portsToScan := makeRange(0, 80)
	timeout := 2

	// Scan duration = timeout x ports to scan (x 2 because we scan tcp and udp)
	fmt.Printf("  Scan maximum duration: %d~ seconds \n", timeout * (len(portsToScan) * 1))
	
	fmt.Println("  Scanning TCP ports")
	tcpScanResult := scanPortList("tcp", host, portsToScan, 1)

	fmt.Println("  Scanning UDP ports")
	udpScanResult := scanPortList("udp", host, portsToScan, 1)

	startTime := time.Now()

	duration := int(time.Since(startTime).Seconds())
	
	

	fmt.Println("  Scan lasted " + strconv.Itoa(duration) +"s")

	fmt.Println("  Scan finished, listing results")
	fmt.Println()

	listScanReport("  tcp/", tcpScanResult)
	listScanReport("  udp/", udpScanResult)

}