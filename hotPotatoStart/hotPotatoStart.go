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
	"github.com/rs/cors"
)

var remotehost string

func enviar(num int) {
	conn, _ := net.Dial("tcp", remotehost)
	defer conn.Close()
	fmt.Fprintf(conn, "%d,%d\n", num, num)
}

func StartKMeans(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	iterations := vars["iterations"]
	nIterations, _ := strconv.Atoi(iterations)
	enviar(nIterations)
	fmt.Fprintf(w, "TOMA TU RESPUESTA en %d iteraciones", nIterations)
}

func main() {
	bufferIn := bufio.NewReader(os.Stdin)
	fmt.Print("Ingrese la IP remota: ")
	ip, _ := bufferIn.ReadString('\n')
	ip = strings.TrimSpace(ip)
	remotehost = fmt.Sprintf("%s:%d", ip, 8002)
	router := mux.NewRouter().StrictSlash(true)
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
	})

	handler := c.Handler(router)
	router.HandleFunc("/kmeans/{iterations}", StartKMeans)
	log.Fatal(http.ListenAndServe(":1234", handler))
}
