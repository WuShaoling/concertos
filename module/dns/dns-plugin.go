package dns

import (
	"sync"
	"log"
)

func (dns *DNSApi) Start() {
	log.Println("start dns module success!")
}

type DNSApi struct {
}

var DnsApi *DNSApi
var once sync.Once

func GetDNSApi() *DNSApi {

	once.Do(func() {
		DnsApi = new(DNSApi)
	})

	return DnsApi
}
