package dhv

import (
	"fmt"
	"io/ioutil"
	"strings"

	gonet "github.com/THREATINT/go-net"
)

type Options struct {
	Hosts             string
	Silent            bool
	Verbose           bool
}


func Process(options *Options) {	
	var found []string

	if options.Verbose {
		fmt.Printf("[*] Loading file: %s\n",options.Hosts)
	}
	bytesRead, _ := ioutil.ReadFile(options.Hosts)
	file_content := string(bytesRead)
	h := strings.Split(file_content, "\n")
	lines := sliceUniqueElements(h)

	if options.Verbose {
		fmt.Printf("[*] Target hosts loaded: %d\n",len(lines))
	}
	
	for _,item := range lines {
		if len(item) > 2 {
			if strings.Contains(item,"*.") {
				a := strings.ReplaceAll(item, "*.", "")
				item = a
			}

			if gonet.IsFQDN(item) {
				if sliceContainsElement(found, item) == false {
					found = append(found, item)
					fmt.Println(item)
				}
			} else {
				if gonet.IsDomain(item) {
					if sliceContainsElement(found, item) == false {
						found = append(found, item)
						fmt.Println(item)
					}
				}
			}
		}
	}
	
}