package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

var remotehost string

func wait() string {
	fmt.Printf("ESTOY ESPERANDO A QUE ME LLAME LA ULTIMA VM")
	direccionIP_API := localAddress()
	hostlocal := fmt.Sprintf("%s:%d", direccionIP_API, 8002)
	ln, _ := net.Listen("tcp", hostlocal)
	defer ln.Close()
	conn, _ := ln.Accept()
	bufferIn := bufio.NewReader(conn)
	load, _ := bufferIn.ReadString('\n')
	load = strings.TrimSpace(load)
	fmt.Printf("Llegó la carga al local!!! :) : %s\n", load)
	return load
}

func enviar(num int) {
	direccionIP_API := localAddress()
	fmt.Println(direccionIP_API)
	conn, _ := net.Dial("tcp", remotehost)
	defer conn.Close()
	fmt.Fprintf(conn, "%d,%d,%s\n", num, num, direccionIP_API)
}

func StartKMeans(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	iterations := vars["iterations"]
	nIterations, _ := strconv.Atoi(iterations)
	enviar(nIterations)
	response := wait()
	fmt.Fprintf(w, "%s", response)
}

func main() {
	bufferIn := bufio.NewReader(os.Stdin)
	fmt.Print("Ingrese la IP remota: ")
	ip, _ := bufferIn.ReadString('\n')
	ip = strings.TrimSpace(ip)
	remotehost = fmt.Sprintf("%s:%d", ip, 8002)
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/kmeans/{iterations}", StartKMeans)
	log.Fatal(http.ListenAndServe(":1234", router))
}

func localAddress() string {
	ifaces, err := net.Interfaces()
	if err != nil {
		log.Print(fmt.Errorf("localAddress: %v\n", err.Error()))
		return "127.0.0.1"
	}
	for _, oiface := range ifaces {
		if strings.HasPrefix(oiface.Name, "Wi-Fi") {
			addrs, err := oiface.Addrs()
			if err != nil {
				log.Print(fmt.Errorf("localAddress: %v\n", err.Error()))
				continue
			}
			for _, dir := range addrs {
				switch d := dir.(type) {
				case *net.IPNet:
					if strings.HasPrefix(d.IP.String(), "192") {
						return d.IP.String()
					}

				}
			}
		}
	}
	return "127.0.0.1"
}
