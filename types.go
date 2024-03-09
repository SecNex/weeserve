package main

import "git.serverify.de/secnex/weeserve/proxy"

type WebServer struct {
	Host    string                     `json:"host"`
	Port    int                        `json:"port"`
	Reverse bool                       `json:"reverse"`
	Targets []proxy.ReverseProxyTarget `json:"targets"`
}
