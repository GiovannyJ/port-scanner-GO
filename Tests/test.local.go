package tests
import (
	"fmt"
	"sync"
	sc "port-scanner-GO/Scanner"
	s "port-scanner-GO/Models"
)

type ScanResult = s.ScanResult

func scanRoutine(r chan string, hostname string, wg *sync.WaitGroup) ([]ScanResult){
	defer wg.Done()
	data, time, err := sc.TimedScan(hostname)
	if err != nil{
		return nil
	}
	r <- time
	return data
}

func Test(){
	var wg sync.WaitGroup
	
	wg.Add(3)
	server :=  make(chan string)
	pi := make(chan string)
	proxy := make(chan string)
	
	go scanRoutine(server, "192.168.1.21", &wg)
	go scanRoutine(pi, "192.168.1.20", &wg)
	go scanRoutine(proxy, "192.168.1.8", &wg)
	
	for i:=0; i<3; i++{
		select{
			case x := <- server:
				fmt.Println("Server Done", x)
			case y := <- pi:
				fmt.Println("PI Done", y)
			case z:= <- proxy:
				fmt.Println("Proxy Done", z)
		}
	}
	wg.Wait()
}