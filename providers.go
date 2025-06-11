package main

type ip_provider_response interface {
	parse(res string) string
}

type ip_provider struct {
	url    string
	method string
	ipv    int8
	opt    map[string]string
	res    ip_provider_response
}

// IPIFY ----------------------------------------------------------------------
type IPIFY struct {
	Ip string
}

func (p IPIFY) parse(res string) string {
	return res
}

// IFCFG ----------------------------------------------------------------------
type IFCFG struct{}

func (p IFCFG) parse(res string) string {
	return res
}
