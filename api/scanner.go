package api

import (
	"net"
	"strconv"
	"time"
	//"fmt"

	"github.com/Sniper10754/snipermapper/api/descriptor"
)

type NetworkProtocol string

const TCP NetworkProtocol = "tcp"
const UDP NetworkProtocol = "udp"

type ScanResult struct {
	State bool
	Port int
	Description descriptor.PortDescription
}

type PortScanner interface {
	ScanPort(port int) ScanResult 
}

type BasicPortScanner struct {
	Address string
}

type TCPPortScanner struct {
	BasicPortScanner
}

type UDPPortScanner struct {
	BasicPortScanner
}

func makeRange(min, max int) []int {
    a := make([]int, max-min+1)
    for i := range a {
        a[i] = min + i
    }
    return a
}

// * Not sthealty 
// 
// ? UDP Scan requires to send a packet to an ip, and then
// ? check if a response is sent.

func (ups UDPPortScanner) ScanPort(port int) ScanResult {
	return ScanResult{State: false, Port: port}
}

func (tps TCPPortScanner) ScanPort(port int) ScanResult {
	address := tps.Address + ":" + strconv.Itoa(port)
	conn, err := net.DialTimeout(string(TCP), address, 60*time.Second)

	if err == nil {
		defer conn.Close()
	}

	return ScanResult{Port: port, State: err == nil}
}

func ScanPorts(scanner PortScanner, start int, finish int, ch chan ScanResult) []ScanResult {
	ports := makeRange(start, finish)
	results := make([]ScanResult, finish)

	
	for _, port := range ports {
		result := scanner.ScanPort(port)

		if ch != nil {
			ch <- result
		}	

		results = append(results, result)
	}

	close(ch)

	return results
}

func NewScanner(protocol NetworkProtocol, address string) PortScanner {
	switch protocol {
	case UDP:
		return UDPPortScanner{BasicPortScanner{Address: address}} 
	
	case TCP:
		return TCPPortScanner{BasicPortScanner{Address: address}}

	default:
		return nil
	}
}