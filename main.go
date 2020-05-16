package main

import (
	"flag"
	"net/http"
	"os"

	"github.com/garrettsparks/ipLookup"
)

const defaultPort = "6060"

func main() {
	listenPort := os.Getenv("PORT")
	if listenPort == "" {
		listenPort = defaultPort
	}

	flag.StringVar(&listenPort, "port", listenPort, "the port to listen on")
	flag.StringVar(&listenPort, "p", listenPort, "the port to listen on")
	flag.Parse()

	http.HandleFunc("/", lookupIP)
	http.ListenAndServe(":"+listenPort, nil)
}

func lookupIP(w http.ResponseWriter, req *http.Request) {
	lookup := ipLookup.New().WithAWS().WithAPIfy().WithWTFIsMyIP().WithLocal()
	ip, err := lookup.GetIP()
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		w.Write([]byte("0.0.0.0"))
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(ip.String()))
}
