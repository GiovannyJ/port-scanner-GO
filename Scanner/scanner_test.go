package scanner

import (
		"testing"
		s "port-scanner-GO/Models"	
	)

type tests = s.TestStruct

func TestScanPort(t *testing.T){
	tests := []tests{
		{Input: 22, Expected: "Closed"},
		{Input: 80, Expected: "Open"},
		{Input: 443, Expected: "Open"}}

	for _, testIN := range tests{
		if output := scanPort("tcp", "google.com", testIN.Input, ""); output.State != testIN.Expected{
			t.Error("Test Failed: input:",testIN.Input ,", expected: ", testIN.Expected,", recieved:", output.State )
		}
	}
}