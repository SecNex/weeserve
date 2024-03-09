package proxy

type ReverseProxyTarget struct {
	Domain string `json:"domain"`
	Host   string `json:"host"`
	Port   int    `json:"port"`
}

func NewReverseProxyTarget(domain string, host string, port int) ReverseProxyTarget {
	return ReverseProxyTarget{Domain: domain, Host: host, Port: port}
}
