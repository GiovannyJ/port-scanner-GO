package parser

import (
	"encoding/json"
	"fmt"
	"os"
	s "packet-sniffing-GO/Models"
)

type Port = s.Port
type ScanResult = s.ScanResult

func Parse(filename string) []Port{
	file, err := os.Open(filename)

	if err != nil{
		fmt.Println(err)
		return nil
	}

	defer file.Close()

	var ports []Port
	json.NewDecoder(file).Decode(&ports)
	
	portList := make([]Port, len(ports))
	// for i, port := range ports{
	// 	portList[i] = port
	// }
	copy(portList, ports)
	return portList
}

func PrettyPrint(data []ScanResult){
	fmt.Println("PORT	STATE	SERVICE")
	for item := range data{
		fmt.Println(data[item].Port,"	", data[item].State,"", data[item].Service)
	}	
}