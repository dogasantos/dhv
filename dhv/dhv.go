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
	Fqdn              bool
	Domain            bool
	SubDomain         bool
	Suffix            bool
	Protocol          bool
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
					parsed := ParseTokens(candidate)
					d := ParseUrlTokens(parsed.String())
					d.Subdomain = parsed.TRD
					d.Domain = parsed.SLD
					d.Tld = parsed.TLD
					
				
					if options.Fqdn == true {
						fmt.Println(parsed.String())
					} else {
						if options.Protocol {
							fmt.Printf("%s ",d.Protocol)
						}
						if options.SubDomain {
							fmt.Printf("%s ",d.Subdomain)
						}
						if options.Domain {
							fmt.Printf("%s.%s",d.Domain,d.Tld)
						}
						if options.Suffix {
							fmt.Printf("%s ",d.Tld)
						}
						fmt.Printf("\n")
					}

				} else {
					if gonet.IsDomain(candidate) {
						parsed := ParseTokens(candidate)
						d := ParseUrlTokens(parsed.String())
						d.Subdomain = parsed.TRD
						d.Domain = parsed.SLD
						d.Tld = parsed.TLD
						
						if options.Fqdn == true {
							fmt.Println(parsed.String())
						} else {
							if options.Protocol {
								fmt.Printf("%s ",d.Protocol)
							}
							if options.SubDomain {
								fmt.Printf("%s ",d.Subdomain)
							}
							if options.Domain {
								fmt.Printf("%s.%s",d.Domain,d.Tld)
							}
							if options.Suffix {
								fmt.Printf("%s ",d.Tld)
							}
							fmt.Printf("\n")
						}
					}
				}
			}(item)
		}
	}
	wg.Wait()	
}