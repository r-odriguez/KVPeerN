package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"net/netip"
	"net/url"
)

func get_self_ipv6(providers []ip_provider) netip.Addr {
	ip, err := get_self_public_ip(providers, 6)

	if err != nil {
		log.Println("Error retrieving public IPv6 from provider!", err)
	}

	return ip
}

func get_self_ipv4(providers []ip_provider) netip.Addr {
	ip, err := get_self_public_ip(providers, 4)

	if err != nil {
		log.Println("Error retrieving public IPv4 from provider!", err)
	}

	return ip
}

func get_self_public_ip(providers []ip_provider, ipv int8) (netip.Addr, error) {
	var client http.Client
	var encoded_url string
	var err error
	var res *http.Response
	var req *http.Request
	var query url.Values = url.Values{}

	for _, p := range providers {
		if p.ipv == ipv {
			for k, v := range p.opt {
				query.Add(k, v)
			}

			client = http.Client{}
			encoded_url = fmt.Sprintf("%s?%s", p.url, query.Encode())
			req, err = http.NewRequest(p.method, encoded_url, nil)

			if err != nil {
				log.Println(err)
			}

			res, err = client.Do(req)

			if err != nil {
				log.Println(err)

				var netError net.Error
				if errors.As(err, &netError) && netError.Timeout() {
					log.Println("Provider didn't respond in time. Error:", netError)
				}
				continue
			}

			body, err := io.ReadAll(res.Body)

			if err != nil {
				log.Println(err)
			}

			if res.StatusCode > 299 {
				log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
			}

			res.Body.Close()
			return netip.ParseAddr(p.res.parse(string(body)))
		}
	}

	return netip.ParseAddr("")
}
