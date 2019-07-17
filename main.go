package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"time"
)

var (
	responseString = "Hello World\n"
	ip, port       string
	server         bool
)

func main() {
	flag.BoolVar(&server, "server", false, "Flag true if this should be started as a server")
	flag.StringVar(&ip, "address", getOutboundIP(), "server name or IP address to connect on/to")
	flag.StringVar(&port, "port", ":8080", "port to connect on/to")
	flag.Parse()
	if port[0] != ':' {
		port = fmt.Sprintf(":%s", port)
	}
	if server {
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, responseString)
		})
		log.Printf("Server listening at http://%v%v", ip, port)
		log.Fatal(http.ListenAndServe(port, nil))
	} else {
		time.Sleep(2 * time.Second)
		url := fmt.Sprintf("http://%v%v", ip, port)
		r, err := http.Get(url)
		if err != nil {
			log.Fatalf("Unable to connect to %s", url)
		}
		bd, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatal(err)
		}
		if string(bd) != responseString {
			log.Fatalf("Unable to connect to %s", url)
		}
		log.Println(string(bd))
	}
}

func getOutboundIP() string {
	conn, err := net.Dial("udp", "8.8.8.8:8080")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return fmt.Sprintf("%v", localAddr.IP)
}
