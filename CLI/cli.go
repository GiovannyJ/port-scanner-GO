package cli

import (
	"fmt"
	"log"
	"os"
	
	"github.com/urfave/cli"

	s "port-scanner-GO/Models"
	p "port-scanner-GO/Parser"
	sc "port-scanner-GO/Scanner"
)

type ScanResult = s.ScanResult

func flagRoutine(c *cli.Context, flag string) error{
	sp, err := sc.Scan(flag)
	
	if err != nil{
		return err
	}
	
	p.PrettyPrint(sp, flag)
	return nil
}

func timeFlagRoutine(c *cli.Context, flag string) error{
	sp, time, err := sc.TimedScan(c.String("t"))
						
	if err != nil {
		fmt.Println("Time: ",time)
		return err
	}
	
	p.PrettyPrint(sp, c.String("t"))
	fmt.Println("Time:",time)
	return nil
}


func CLI(){
	app := cli.NewApp()
	app.Name = "Port Scanner GO"
	app.Usage = "Scan Ports on a network with TCP or UDP"

	flags := []cli.Flag{
		cli.StringFlag{
			Name: "s",
			Value: "google.com",
			Usage: "Scan a list of ports",
		},
		cli.StringFlag{
			Name: "t",
			Value: "google.com",
			Usage: "Scan a list of ports and show the time it took",

		},
	}
	
	app.Commands = []cli.Command{
		{
			Name: "scan",
			Usage: "Scanning Ports",
			Flags: flags,
			Action: func (c *cli.Context) error {
				if c.NumFlags() > 0{
					if !c.IsSet("t") && c.IsSet("s") {						
						flagRoutine(c, c.String("s"))
						
					}else if !c.IsSet("s") && c.IsSet("t") {
						timeFlagRoutine(c, c.String("t"))

					}else{
						log.Fatal("Only Use One Flag (-s -t)")
						return nil
					}
				}
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil{
		log.Fatal(err)
	}
}
