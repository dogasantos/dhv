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
 
func worker(id int, jobs <-chan string, results chan<-string) {
	var found []string
	
	for item := range jobs {
		if gonet.IsFQDN(item) {
			if sliceContainsElement(found, item) == false {
				found = append(found, item)
				fmt.Println(item)
				results <- item
			}
		} else {
			if gonet.IsDomain(item) {
				if sliceContainsElement(found, item) == false {
					found = append(found, item)
					fmt.Println(item)
					results <- item
				}
			}
		}
	}
}



func Process(options *Options) {	
	var found []string
	const numJobs = 10

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
	
	jobs := make(chan string, numJobs)
	results := make(chan string, numJobs)

	for w := 1; w <= 100; w++ {
		go worker(w, jobs, results)
	}

	for _, item := range lines {
		if len(item) > 2 {
			jobs <- item
		}
	}
	close(jobs)

	if options.Verbose {
		fmt.Printf("[*] Found %d valid entries\n",len(found))
	}
}