package dhv

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	gonet "github.com/THREATINT/go-net"
)

type Options struct {
	Hosts             string
	OutputFile        string
	Silent            bool
	Verbose           bool
}

func Process(options *Options) {	
	
	var found []string
	var hostname string 

	bytesRead, _ := ioutil.ReadFile(options.Hosts)
	file_content := string(bytesRead)
	h := strings.Split(file_content, "\n")
	lines := sliceUniqueElements(h)

	
	if options.Verbose {
		fmt.Printf("[*] Target hosts loaded: %d\n",len(lines))
	}
	
	for _, item := range lines {
		if len(item) > 2 {
			if gonet.IsFQDN(item) {
				if sliceContainsElement(found, item) == false {
					found = append(found, hostname)
				}
			} else {
				if gonet.IsDomain(item) {
					if sliceContainsElement(found, item) == false {
						found = append(found, hostname)
					}
				}
			}
		}
	}
		

	fmt.Printf("[*] Found %d hosts\n",len(found))

	if len(options.OutputFile) >0 {
		file, _ := os.Create(options.OutputFile)
		writer := bufio.NewWriter(file)
		for _, fh := range found {
			_, _ = writer.WriteString(fh + "\n")
		}
		writer.Flush()
	}
	for _, fh := range found {
		fmt.Println(fh)
	}
}