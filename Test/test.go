package test

import (
	"testing"
	"github.com/stretchr/testify/assert"
	sc "port-scanner-GO/Scanner"
	s "port-scanner-GO/Models"
)

type ScanResult = s.ScanResult

func Test(t *testing.T){
	var known ScanResult
	known.Port = "80"
	known.Service = "http"
	known.State = "Open"

	out := sc.Scan("shnybones.duckdns.org")
	assert.Equal(t, known, out)
}