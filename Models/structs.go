package structs


type  ScanResult struct{
	Port string
	State string
	Service string
}

type Port struct{
	Port int `json:"port"`
	Service string `json:"service"`
}
