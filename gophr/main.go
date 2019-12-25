package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/jayvib/golog"
)

var port string
var host string
var debug bool

func init() {
	flag.StringVar(&port, "port", "8080", "Port of the application")
	flag.StringVar(&host, "host", "127.0.0.1", "Host of the application")
	flag.BoolVar(&debug, "debug", false, "Debug mode")

	flag.Parse()

	if debug {
		golog.Warning("DEBUGGING MODE!")
		golog.SetLevel(golog.DebugLevel)
	}
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))
	addr := fmt.Sprintf("%s:%s", host, port)
	golog.Info("Listening in ", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatal(err)
	}
}
