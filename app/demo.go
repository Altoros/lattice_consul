// This web app balances between two mysql instances,
// discovered using consul

package main

import (
	"flag"
	"fmt"
	"github.com/benschw/dns-clb-go/clb"
	"log"
	"net/http"
)

// Returns the svcName service's ip address and port
// svcName like `mysql`, without `.service.consul`
// Assumes consul agent is running locally on port 8600

func getAddress(svcName string) (string, error) {
	c := clb.NewClb("127.0.0.1", "8600", clb.Random)

	srvRecord := svcName + ".service.consul"
	address, err := c.GetAddress(srvRecord)
	if err != nil {
		return "", err
	}

	return address.String(), nil
}

// Func prints mysql service ip address, port and
// hostname of the node it runs on

func demo(w http.ResponseWriter, req *http.Request) {
	addStr, err := getAddress("mysql")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(w, "Discovered Address: '%s'\n", addStr)

	// TODO: print node hostname
}

var addr = flag.String("addr", ":8080", "http service address")

func main() {
	flag.Parse()
	http.Handle("/", http.HandlerFunc(demo))

	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
