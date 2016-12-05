package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

var (
	host    string
	port    int
	verbose bool
)

func main() {
	flag.StringVar(&host, "h", "", "hostname.")
	flag.IntVar(&port, "p", 8600, "port number.")
	flag.BoolVar(&verbose, "v", false, "verbose mode.")
	flag.Parse()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "hello, world")
	})

	addr := fmt.Sprintf("%s:%d", host, port)
	log.Printf("listen %s", addr)

	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
