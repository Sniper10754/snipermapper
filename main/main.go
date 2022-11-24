package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"syscall"
	"time"

	snipermapper "github.com/Sniper10754/snipermapper/api"
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

func listScanReport(prefix string, results []snipermapper.ScanResult) {
	for i := 0; i < len(results); i++ {
		scan := results[i]

		if scan.State {
			fmt.Printf("%s Port open: %d\n", prefix, scan.Port)
		}
	}
	
}

func scanPortsWithPrompt(scanner snipermapper.PortScanner, start int, finish int) []snipermapper.ScanResult {
	ch := make(chan snipermapper.ScanResult)
	var ports []snipermapper.ScanResult
	
	go func() {
		ports = snipermapper.ScanPorts(scanner, start, finish, ch)
	}()

	for {
		val, ok := (<- ch)

		if ok {
			fmt.Printf("%d -> %s\r", val.Port, strconv.FormatBool(val.State))
		} else {
			break
		}
	}

	return ports
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func main() {
	stealth := contains(os.Args, "-stealth")
	fmt.Println(title)

	if !stealth {
		fmt.Println("Stealth scan enabled")
	}

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

	tcpPortScanner := snipermapper.NewScanner(snipermapper.TCP, host)
	var udpPortScanner snipermapper.PortScanner
	
	if !stealth {
		udpPortScanner = snipermapper.NewScanner(snipermapper.UDP, host)
	}
	// Scan duration = timeout x ports to scan (x 2 because we scan tcp and udp)
	fmt.Printf("  Scan maximum duration: %d~ seconds \n", timeout * (portFinish * 2))
	
	fmt.Println("  Scanning TCP ports")
	tcpScanResult := scanPortsWithPrompt(tcpPortScanner, portStart, portFinish)
	
	var udpScanResult []snipermapper.ScanResult = make([]snipermapper.ScanResult, portFinish)

	if !stealth {
		fmt.Println("  Scanning UDP ports")
		udpScanResult = scanPortsWithPrompt(udpPortScanner, portStart, portFinish)
	}

	startTime := time.Now()

	duration := int(time.Since(startTime).Seconds())

	fmt.Println("  Scan lasted " + strconv.Itoa(duration) +"s")

	fmt.Println("  Scan finished, listing results")
	fmt.Println()

	listScanReport("  tcp/", tcpScanResult)
	listScanReport("  udp/", udpScanResult)

}