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
	
	flag.StringVar(&options.Hosts,                  "L", "", "File input with list of hosts")    
	flag.BoolVar(&options.Fqdn,                   	"fqdn", true, "show the entire fqdn with no parsing at all (default)")      
	flag.BoolVar(&options.Domain,                   "domain", false, "show domain portion")      
	flag.BoolVar(&options.SubDomain,                "subdomain", false, "show subdomain portion")
	flag.BoolVar(&options.Suffix,                   "suffix", false, "show suffix portion")      
	flag.BoolVar(&options.Protocol,                 "protocol", false, "show protocol portion")  

	flag.Parse()

	flag.Usage = func() {
		fmt.Printf("Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
	}

	if flag.NFlag() == 0 {
		flag.Usage()
		os.Exit(1)
	}
	if options.Domain == true {
		options.Fqdn = false
	}
	if options.SubDomain == true {
		options.Fqdn = false
	}
	if options.Suffix == true {
		options.Fqdn = false
	}
	if options.Protocol == true {
		options.Fqdn = false
	}
	dhv.Process(&options)

}