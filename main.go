package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type WebServer struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

func getFlags() WebServer {
	host := flag.String("host", "localhost", "host")
	port := flag.Int("port", 8080, "port")
	flag.Parse()
	return WebServer{Host: *host, Port: *port}
}

func (w WebServer) Run() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ip := r.RemoteAddr
		if idx := strings.LastIndex(ip, ":"); idx != -1 {
			ip = ip[:idx]
		}
		log.Println("Anfrage von IP-Adresse:", ip) // Protokolliere die IP-Adresse in der Konsole
		fmt.Fprintf(w, "Hello World! Deine IP-Adresse ist %s", ip)
	})

	addr := fmt.Sprintf("%s:%d", w.Host, w.Port)
	fmt.Printf("Server is running on http://%s:%d\n", w.Host, w.Port)
	http.ListenAndServe(addr, nil)
}

func main() {
	server := getFlags()
	go server.Run() // Starte den Webserver in einer Go-Routine

	// Hauptschleife
	fmt.Println("Press Ctrl+C to exit!")
	select {}
}
