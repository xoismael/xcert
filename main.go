package main

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"
)

func prep_input(user_input string) (output string) {
	pattern := `^([0-9a-zA-Z-.]+)+\:?([0-9]{1,5})$`

	rp := regexp.MustCompile(pattern)
	groups := rp.FindAllStringSubmatch(user_input, -1)
	if len(groups) > 0 {
		for _, group := range groups {
			hostname := group[1]
			port := group[2]
			if len(strings.TrimSpace(port)) == 0 {
				fmt.Println("No port specified")
				break
			} else {
				switch port {
				case "443":
					output = fmt.Sprintf("https://%s", hostname)
				case "80":
					output = fmt.Sprintf("http://%s", hostname)
				default:
					output = fmt.Sprintf("https://%s:%s", hostname, port)
				}
			}

		}

	} else {
		fmt.Println("Input should be in the form of <host:port>")
	}
	return

}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		url := prep_input(sc.Text())
		client := &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true,
				},
			},
			Timeout: time.Second * 10,
		}
		resp, err := client.Get(url)
		if err != nil {
			continue
		}
		if resp.TLS == nil {
			continue
		}
		cert := resp.TLS.PeerCertificates[0]
		fmt.Printf("%s\t%s %s\n", url, cert.Subject.CommonName, cert.Subject.Organization)
		defer resp.Body.Close()
	}
}
