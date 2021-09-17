package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/dogasantos/dhv/dhv"
)

const banner = `
FQDN VALIDATOR
==============
`
const Version = `0.1`

func showBanner() {
	fmt.Printf("%s", banner)
	fmt.Printf("\t\t\t\t\t\t\t\tversion: %s\n\n",Version)
}

func main() {	
	
	options := dhv.Options{}
	
	flag.StringVar(&options.Hosts, 			"L", "", "File input with list of hosts")
	flag.BoolVar(&options.Verbose, 			"verbose", false, "Verbose output")

	flag.Parse()

	flag.Usage = func() {
		fmt.Printf("Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
	}

	if flag.NFlag() == 0 {
		flag.Usage()
		os.Exit(1)
	}

	dhv.Process(&options)

}