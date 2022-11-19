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

func ScanPort(netprotocol NetworkProtocol, host string, port int, timeout int32) bool {
	address := host + ":" + strconv.Itoa(port) 
	conn, err := net.DialTimeout(string(netprotocol), address, time.Duration(time.Duration(timeout).Seconds()))


	if err == nil {defer conn.Close()}
	
	return err == nil
}

func ScanPorts(netprotocol NetworkProtocol, host string, ports []int, timeout int32) []ScanResult {
	scanResults := make([]ScanResult, len(ports))
	
	for i := 0; i < len(scanResults); i++ {
		scanResults[i] = ScanResult{ScanPort(netprotocol, host, ports[i], timeout), ports[i]}
	}

	return scanResults
}