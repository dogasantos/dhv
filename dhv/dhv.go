package dhv

import (
	"fmt"
	"io/ioutil"
	"strings"
	"sync"

	gonet "github.com/THREATINT/go-net"
)

type Options struct {
	Hosts             string
	Silent            bool
	Verbose           bool
}



func Process(options *Options) {
	var wg sync.WaitGroup
	
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
			wg.Add(1)
			go func(candidate string) {
				defer wg.Done()
				if gonet.IsFQDN(candidate) {
					fmt.Println(candidate)
				} else {
					if gonet.IsDomain(candidate) {
						fmt.Println(candidate)
					}
				}
			}(item)
		}
	}
	wg.Wait()	
}