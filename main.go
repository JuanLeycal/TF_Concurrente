package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/sadaakisz/go-api/kmeans"
	"github.com/sadaakisz/go-api/model"
)

var data = [][]string{}
var listaCiudadano = []model.Ciudadano{}
var nCiudadanos []kmeans.Node = []kmeans.Node{}
var remotehost string

func readCsv(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file: "+filePath, err)
	}
	defer f.Close()
	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for: "+filePath, err)
	}
	return records
}

func populateData() {
	for i, row := range data {
		a, _ := strconv.Atoi(row[0])
		b, _ := strconv.Atoi(row[1])
		c, _ := strconv.Atoi(row[2])
		d, _ := strconv.Atoi(row[3])
		e, _ := strconv.Atoi(row[4])
		object := model.Ciudadano{
			ID:                     i,
			ESTRATO_SOCIOECONOMICO: float64(a),
			SEGURIDAD_NOCTURNA:     float64(b),
			GRUPOS_EDAD:            float64(c),
			CONFIANZA_POLICIA:      float64(d),
			PRONTO_DELITO:          float64(e),
		}
		listaCiudadano = append(listaCiudadano, object)
		nNode := kmeans.Node{}
		nNode = append(nNode, float64(a), float64(b), float64(c), float64(d), float64(e))
		nCiudadanos = append(nCiudadanos, nNode)
	}
	listaCiudadano = listaCiudadano[1:]
	nCiudadanos = nCiudadanos[1:]
}

func runKMeans(kClusters int) {
	if success, nList, nCentroids, maxIter := kmeans.KMeansInit(nCiudadanos, kClusters, 50); success {

		encodedStr := kmeans.TrainingEncode(nList, nCentroids, maxIter)
		enviar(encodedStr)
		toDecodeStr := wait()
		nList, nCentroids, maxIter = kmeans.TrainingDecode(toDecodeStr)

		//esta logica deberia ir en el listener del start
		fmt.Println("Centroids:")
		for _, centroid := range nCentroids {
			fmt.Println(centroid)
		}
		for i, ciudadano := range nCiudadanos {
			clusterI := kmeans.Nearest(ciudadano, nCentroids)
			listaCiudadano[i].CLUSTER = clusterI + 1
			fmt.Println(ciudadano, "Cluster:", clusterI+1)
		}
	}
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	stkClusters := vars["kClusters"]
	kClusters, _ := strconv.Atoi(stkClusters)
	if kClusters == 0 {
		kClusters = 4
	}
	runKMeans(kClusters)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(listaCiudadano)
}

//Distribution portion start

func wait() string {
	//fmt.Printf("ESTOY ESPERANDO A QUE ME LLAME LA ULTIMA VM")
	direccionIP_API := localAddress()
	hostlocal := fmt.Sprintf("%s:%d", direccionIP_API, 8002)
	ln, _ := net.Listen("tcp", hostlocal)
	defer ln.Close()
	conn, _ := ln.Accept()
	bufferIn := bufio.NewReader(conn)
	load, _ := bufferIn.ReadString('\n')
	load = strings.TrimSpace(load)
	//fmt.Printf("Llegó la carga al local!!! :) : %s\n", load)
	return load
}

func enviar(encodedStr string) {
	direccionIP_API := localAddress()
	fmt.Println(direccionIP_API)

	conn, _ := net.Dial("tcp", remotehost)
	defer conn.Close()
	fmt.Fprintf(conn, "%s/%s/%d/%d", encodedStr, direccionIP_API, 50, 50)
}

func localAddress() string {
	ifaces, err := net.Interfaces()
	if err != nil {
		log.Print(fmt.Errorf("localAddress: %v\n", err.Error()))
		return "127.0.0.1"
	}
	for _, oiface := range ifaces {
		//Change to Ethernet if you have cable
		if strings.HasPrefix(oiface.Name, "Ethernet") {
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

//Distribution portion end

func main() {
	fmt.Println("Running")
	r := mux.NewRouter()
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
	})
	handler := c.Handler(r)

	data = readCsv("./DatasetSelectivo.csv")

	populateData()

	bufferIn := bufio.NewReader(os.Stdin)
	fmt.Print("Ingrese la IP remota: ")
	ip, _ := bufferIn.ReadString('\n')
	ip = strings.TrimSpace(ip)
	remotehost = fmt.Sprintf("%s:%d", ip, 8002)

	r.HandleFunc("/{kClusters}", HomeHandler)
	log.Fatal(http.ListenAndServe(":3000", handler))
}
