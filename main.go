package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
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
	if success, nCentroids := kmeans.KMeansInit(nCiudadanos, kClusters, 50); success {
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
