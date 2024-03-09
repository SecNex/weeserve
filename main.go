package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"git.serverify.de/secnex/weeserve/proxy"
)

func (w WebServer) Run() {
	if w.Reverse {
		w.startReverseProxy()
	} else {
		w.startServer()
	}
}

func (x WebServer) startServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ip := r.RemoteAddr
		if idx := strings.LastIndex(ip, ":"); idx != -1 {
			ip = ip[:idx]
		}
		method := r.Method
		uri := r.RequestURI
		log.Printf("%s %s %s", ip, method, uri) // Logge die Anfrage
		fmt.Fprintf(w, "Hello World!")          // Sende "Hello World!" als Antwort
	})

	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Du hast die /test Route aufgerufen!")
	})

	http.HandleFunc("/ip", func(w http.ResponseWriter, r *http.Request) {
		ip := r.RemoteAddr
		if idx := strings.LastIndex(ip, ":"); idx != -1 {
			ip = ip[:idx]
		}
		fmt.Fprintf(w, "Deine IP-Adresse ist: %s", ip)
	})

	addr := fmt.Sprintf("%s:%d", x.Host, x.Port)
	fmt.Printf("Server is running on http://%s:%d\n", x.Host, x.Port)
	http.ListenAndServe(addr, nil)
}

func (x WebServer) startReverseProxy() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ip := r.RemoteAddr
		if idx := strings.LastIndex(ip, ":"); idx != -1 {
			ip = ip[:idx]
		}
		method := r.Method
		path := r.URL.Path
		calledDNS := r.Host
		log.Printf("%s %s %s%s", ip, method, calledDNS, path) // Logge die Anfrage
		// Check if the request is for a known domain
		for _, target := range x.Targets {
			if target.Domain == calledDNS {
				// Send response
				fmt.Fprintf(w, "Reverse Proxy: %s:%d", target.Host, target.Port)
			} else {
				// Send 404
				http.Error(w, "404 Not Found", http.StatusNotFound)
			}
		}
	})
	addr := fmt.Sprintf("%s:%d", x.Host, x.Port)
	fmt.Printf("Reverse Proxy is running on http://%s:%d\n", x.Host, x.Port)
	http.ListenAndServe(addr, nil)
}

func main() {
	server := GetFlags()
	if server.Reverse {
		// Reverse Proxy
		server.AddReverseProxyTarget(proxy.NewReverseProxyTarget("localhost:8080", "serverify.de", 443))
	}
	go server.Run() // Starte den Webserver in einer Go-Routine

	// Hauptschleife
	fmt.Println("Press Ctrl+C to exit!")
	select {}
}
