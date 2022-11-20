package api

import (
	"net"
	"strconv"
	"time"
)

type NetworkProtocol string

type ScanResult struct {
	State bool
	Port int
}

func ScanPort(netprotocol NetworkProtocol, host string, port int, timeout int32) ScanResult {
	address := host + ":" + strconv.Itoa(port) 
	conn, err := net.DialTimeout(string(netprotocol), address, time.Duration(time.Duration(timeout).Seconds()))


	if err == nil {defer conn.Close()}
	
	return ScanResult{err == nil, port}
}

// ScanPorts
// 
// Allows to scan a range of ports contained in ports argument.
//
// netprotocol: NetworkProtocol
// Can be "tcp" or "udp"
//
// host: string
// Host to scan
//
// timeout: int32
// Timeoout to use for each port scan, the lower the faster the scan is
//
// ch: chan ScanResult
// A channel which is notified about every port scanned
// Enables to know the scan state.
func ScanPorts(netprotocol NetworkProtocol, host string, ports []int, timeout int32, ch chan ScanResult) []ScanResult {
	scanResults := make([]ScanResult, len(ports))
	
	for i := 0; i < len(scanResults); i++ {
		scanResults[i] = ScanPort(netprotocol, host, ports[i], timeout)

		// Allow to know scan state/point
		if ch != nil{
			ch <- scanResults[i]
		}
	}

	return scanResults
}