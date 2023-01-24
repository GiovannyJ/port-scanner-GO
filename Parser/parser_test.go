package parser

import (
	"testing"
)



func TestParse(t *testing.T){

	if result := Parse("../ports.short.json"); result[0].Port != 20 ||  result[0].Service != "FTP data transfer" {
		t.Error("Test Failed:  recieved:", result[0].Port, " " ,result[0].Service )
	}

}