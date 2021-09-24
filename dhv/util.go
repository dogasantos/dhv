package dhv

import (
	"github.com/bobesa/go-domain-util/domainutil"
	"github.com/weppos/publicsuffix-go/publicsuffix"
)

type DomainTokens struct {
    Protocol string
	Subdomain string
	Domain string
	Tld string
}

func ParseUrlTokens(value string) (*DomainTokens){
	var d DomainTokens
    d.Protocol = domainutil.Protocol(value)
	d.Subdomain = domainutil.Subdomain(value)
	d.Domain = domainutil.Domain(value)
	d.Tld = domainutil.DomainSuffix(value)
	return &d
}

func ParseTokens(value string) (*publicsuffix.DomainName, error) {
	var options publicsuffix.ParserOption
	options.PrivateDomains = false
	options.ASCIIEncoded = false
	parsed,err := publicsuffix.Parse(value)
	return parsed,err

}

func sliceContainsElement(slice []string, element string) bool {
	retval := false
	for _, e := range slice {
		if e == element {
			retval = true
		}
	}
	return retval
}



func sliceUniqueElements(slice []string) []string {
    keys := make(map[string]bool)
    list := []string{}
    for _, entry := range slice {
        if _, value := keys[entry]; !value {
            keys[entry] = true
            if len(entry) > 2 {
                list = append(list, entry)
            }
        }
    }
    return list
}

func sliceDifference(slice1 []string, slice2 []string) []string {
    var diff []string

    for i := 0; i < 2; i++ {
        for _, s1 := range slice1 {
            found := false
            for _, s2 := range slice2 {
                if s1 == s2 {
                    found = true
                    break
                }
            }
            if !found {
                diff = append(diff, s1)
            }
        }
        if i == 0 {
            slice1, slice2 = slice2, slice1
        }
    }
    return diff
}

