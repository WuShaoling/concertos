package dns

import (
	"sync"
	"net/http"
	"io/ioutil"
	"github.com/concertos/module/common"
	"strings"
)

func (dns *DNSApi) Add(domain, ip string) string {
	url := "http://" + common.GetDNSServer_API_ADDR() + "/domain"
	payload := strings.NewReader("{\"IP\":\"" + ip + "\", \"Domain\":\"" + domain + "\"}")

	req, _ := http.NewRequest("PUT", url, payload)
	req.Header.Add("content-type", "application/json")

	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()

	return res.Status
}

func (dns *DNSApi) Delete(domain string) string {

	url := "http://" + common.GetDNSServer_API_ADDR() + "/domain/" + domain

	req, _ := http.NewRequest("DELETE", url, nil)
	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	return res.Status
}

func (dns *DNSApi) GetAll() string {
	url := "http://" + common.GetDNSServer_API_ADDR() + "/domain"
	req, _ := http.NewRequest("GET", url, nil)
	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	return string(body)
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
