package scanner

import (
	"net"
	s "port-scanner-GO/Models"
	p "port-scanner-GO/Parser"
	"strconv"
	"sync"
	"time"
)

type ScanResult = s.ScanResult
type Port = s.Port


//merge channels together
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

	conn, err := net.DialTimeout(protocol, address, 60*time.Second)

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

	tcpQueue := make(chan ScanResult, 100)
	udpQueue := make(chan ScanResult, 100)

	ports := p.Parse("ports.json")
	
	//! ADJUST AMOUNT OF WORKER DEPENDING ON YOUR SYSTEM
	workerPool := make(chan struct{}, 100)
	
	for _, elem := range ports {
        workerPool <- struct{}{} // acquire a worker
        wg.Add(1)
        go func(port Port) {
            defer wg.Done()
            defer func() { <-workerPool }() // release the worker
            tcpQueue <- scanPort("tcp", hostname, port.Port, port.Service)
            udpQueue <- scanPort("udp", hostname, port.Port, port.Service)
        }(elem)
    }

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