package main

import (
	"flag"

	"git.serverify.de/secnex/weeserve/proxy"
)

func GetFlags() WebServer {
	host := flag.String("host", "localhost", "host")
	port := flag.Int("port", 8080, "port")
	reverse := flag.Bool("reverse", false, "reverse")
	flag.Parse()
	return WebServer{Host: *host, Port: *port, Reverse: *reverse}
}

func NewWebServer(host string, port int, reverse bool) WebServer {
	return WebServer{Host: host, Port: port, Reverse: reverse, Targets: []proxy.ReverseProxyTarget{}}
}

func (w *WebServer) AddReverseProxyTarget(target proxy.ReverseProxyTarget) {
	w.Targets = append(w.Targets, target)
}
