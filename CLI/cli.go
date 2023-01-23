package cli

import (
	// "fmt"
	"fmt"
	"log"
	"os"
	s "port-scanner-GO/Models"
	p "port-scanner-GO/Parser"
	sc "port-scanner-GO/Scanner"

	"github.com/urfave/cli"
)

type ScanResult = s.ScanResult

func CLI(){
	app := cli.NewApp()
	app.Name = "Port Scanner GO"
	app.Usage = "Scan Ports on a network with TCP or UDP"

	flags := []cli.Flag{
		cli.StringFlag{
			Name: "host",
			Value: "google.com",
		},
	}
	
	app.Commands = []cli.Command{
		{
			Name: "SP",
			Usage: "Scanning Ports",
			Flags: flags,
			Action: func (c *cli.Context) error {
				sp, err := sc.Scan(c.String("host"))
				if err != nil{
					return err
				}
				p.PrettyPrint(sp)
				return nil
			},
		},
		{
			Name: "TS",
			Usage: "Scanning Ports and giving time duration",
			Flags: flags,
			Action: func (c *cli.Context)  error{
				sp, time, err := sc.TimedScan(c.String("host"))
				if err != nil {
					fmt.Println("Time: ",time)
					return err
				}
				p.PrettyPrint(sp)
				fmt.Println("Time:",time)
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil{
		log.Fatal(err)
	}
}
