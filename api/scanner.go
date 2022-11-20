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

// ScanPort
// Scans a single port on specified host
func ScanPort(netprotocol NetworkProtocol, host string, port int, timeout int32) ScanResult {
	address := host + ":" + strconv.Itoa(port) 
	conn, err := net.DialTimeout(string(netprotocol), address, time.Duration(time.Duration(timeout).Seconds()))


	if err == nil {defer conn.Close()}
	
	return ScanResult{err == nil, port}
}

/** */
func ScanPorts(netprotocol NetworkProtocol, host string, ports []int, timeout int32, ch *chan ScanResult) []ScanResult {
	scanResults := make([]ScanResult, len(ports))
	
	for i := 0; i < len(scanResults); i++ {
		scanResults[i] = ScanPort(netprotocol, host, ports[i], timeout)

		// Allow to know scan state/point
		if *ch != nil{
			*ch <- scanResults[i]
		}
	}

	return scanResults
}