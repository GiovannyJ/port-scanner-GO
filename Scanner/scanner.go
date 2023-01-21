package scanner

import (
	"net"
	s "packet-sniffing-GO/Models"
	p "packet-sniffing-GO/Parser"
	"strconv"
	"sync"
	"time"
)

type ScanResult = s.ScanResult

func tcpRoutine(c chan ScanResult, hostname string, port int, service string, wg *sync.WaitGroup){
	defer wg.Done()
	result := scanPort("tcp", hostname, port, service)
	c <- result
}

func udpRoutine(c chan ScanResult, hostname string, port int, service string, wg *sync.WaitGroup){
	defer wg.Done()
	result := scanPort("udp", hostname, port, service)
	c <- result
}

func merge(cs ...<-chan ScanResult) <-chan ScanResult {
	out := make(chan ScanResult)
	var wg sync.WaitGroup
	wg.Add(len(cs))
	for _, c := range cs {
		go func(c <-chan ScanResult) {
			for v := range c {
				out <- v
			}
			wg.Done()
		}(c)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}


func scanPort(protocol, hostname string, port int, service string) ScanResult {
	result := ScanResult{Port: protocol + "/" + strconv.Itoa(port)}
	address := hostname + ":" + strconv.Itoa(port)

	conn, err := net.DialTimeout(protocol, address, 15*time.Second)

	if err != nil {
		result.State = "Closed"
		result.Service = service
		return result
	}

	defer conn.Close()
	result.State = "Open"
	result.Service = service
	return result
}


func initScan(hostname string) []ScanResult {
	var wg sync.WaitGroup

	var results []ScanResult

	tcpQueue := make(chan ScanResult, 10)
	udpQueue := make(chan ScanResult, 10)

	ports := p.Parse("ports.json")

	// ports := []int{443,22,80}
	
	for elem := range ports{
		wg.Add(2)
		go tcpRoutine(tcpQueue, hostname, ports[elem].Port, ports[elem].Service, &wg)
		go udpRoutine(udpQueue, hostname, ports[elem].Port, ports[elem].Service, &wg)
	}

	//! asynch
	// for elem := range ports{
	// 	results = append(results, scanPort("tcp", hostname, ports[elem]))
	// 	results = append(results, scanPort("udp", hostname, ports[elem]))
	// }
	// for i := 1; 1<5; i++{
	// 	wg.Add(2)
	// 	go tcpRoutine(tcpQueue, hostname, i, &wg)
	// 	go udpRoutine(udpQueue, hostname, i, &wg)
	// }

	wg.Wait()
	close(tcpQueue)
	close(udpQueue)
	queue := merge(tcpQueue, udpQueue)

	for elem := range queue{
		results = append(results, elem)
	}
	return results
}


func Scan(hostname string) ([]ScanResult, string){
	start := time.Now()
	
	results := initScan(hostname)
	
	fin := time.Since(start)
	p.PrettyPrint(results)

	return results, fin.String()
}