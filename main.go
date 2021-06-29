package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/sadaakisz/go-api/kmeans"
	"github.com/sadaakisz/go-api/model"
)

var data = [][]string{}
var listaCiudadano = []model.Ciudadano{}
var nCiudadanos []kmeans.Node = []kmeans.Node{}

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
			ESTRATO_SOCIOECONOMICO: float32(a),
			SEGURIDAD_NOCTURNA:     float32(b),
			GRUPOS_EDAD:            float32(c),
			CONFIANZA_POLICIA:      float32(d),
			PRONTO_DELITO:          float32(e),
		}
		listaCiudadano = append(listaCiudadano, object)
		nNode := kmeans.Node{}
		nNode = append(nNode, float32(a), float32(b), float32(c), float32(d), float32(e))
		nCiudadanos = append(nCiudadanos, nNode)
	}
	listaCiudadano = listaCiudadano[1:]
	nCiudadanos = nCiudadanos[1:]
}

func runKMeans(kClusters int) {
	if success, nCentroids := kmeans.KMeans(nCiudadanos, kClusters, 50); success {
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
	//fmt.Println(data)

	populateData()

	r.HandleFunc("/{kClusters}", HomeHandler)
	log.Fatal(http.ListenAndServe(":3000", handler))
}
