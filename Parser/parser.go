package parser

import (
	"encoding/json"
	"fmt"
	"os"
	s "port-scanner-GO/Models"
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
	
	return ports
}

func PrettyPrint(data []ScanResult, hostname string){
	fmt.Println("PORT		STATE		SERVICE")
	fmt.Println("===============================================")
	for item := range data{
		fmt.Printf("| %-10v | %-10v | %-10v |\n", data[item].Port, data[item].State, data[item].Service)
		// fmt.Printf("%-10v",data[item].Port,"	", data[item].State,"", data[item].Service)
	}
	fmt.Println("================================================")
	fmt.Println("Scanned:", hostname)
}