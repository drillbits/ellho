package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net"
	"net/http"
	"os"
)

var (
	host    string
	port    int
	verbose bool
)

type VerboseInfo struct {
	Hostname   string
	Interfaces map[string][]net.IP
}

func NewVerboseInfo() (*VerboseInfo, error) {
	info := &VerboseInfo{}

	hostname, err := os.Hostname()
	if err != nil {
		return nil, err
	}
	info.Hostname = hostname

	interfaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}

	ifaces := make(map[string][]net.IP)
	for _, i := range interfaces {
		var ips []net.IP
		if v, ok := ifaces[i.Name]; ok {
			ips = v
		}

		addrs, err := i.Addrs()
		if err != nil {
			return nil, err
		}
		if len(addrs) == 0 {
			continue
		}

		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			ips = append(ips, ip)
		}
		ifaces[i.Name] = ips
	}
	info.Interfaces = ifaces

	return info, nil
}

func main() {
	flag.StringVar(&host, "h", "", "hostname.")
	flag.IntVar(&port, "p", 8600, "port number.")
	flag.BoolVar(&verbose, "v", false, "verbose mode.")
	flag.Parse()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if !verbose {
			fmt.Fprint(w, "hello, world\n")
			return
		}

		t, err := template.ParseFiles("./index.html")
		if err != nil {
			fmt.Fprintf(w, "failed to Parse template: %s\n", err)
			return
		}

		info, err := NewVerboseInfo()
		if err != nil {
			fmt.Fprintf(w, "failed to create verbose info: %s\n", err)
			return
		}

		err = t.Execute(w, info)
		if err != nil {
			fmt.Fprintf(w, "failed to execute the parsed template: %s\n", err)
			return
		}
	})

	addr := fmt.Sprintf("%s:%d", host, port)
	log.Printf("listen %s", addr)

	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
