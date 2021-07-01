package kmeans

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func toString(nList []Node) string {
	var sNode [][]string
	var pog []string

	for i := 0; i < len(nList); i++ {
		var aux []string
		for j := 0; j < len(nList[i]); j++ {
			s := fmt.Sprintf("%f", nList[i][j])
			if j == len(nList[i])-1 {
				aux = append(aux, s+",")
			} else {
				aux = append(aux, s+"-")
			}

		}
		sNode = append(sNode, aux)
	}
	for i := 0; i < len(sNode); i++ {
		pog = append(pog, strings.Join(sNode[i], ""))
	}
	sLong := strings.Join(pog, "")

	return sLong
}
func toInt(value int) string {
	sNum := strconv.Itoa(value)

	return sNum
}

func decodeNode(array []string) (nNode [][]float64) {

	var example [][]string
	for i := 0; i < len(array); i++ {
		s := strings.Split(array[i], "-")
		example = append(example, s)
	}
	for i := 0; i < len(example)-1; i++ {
		var aux []float64
		if len(example[i]) != 0 {
			for j := 0; j < len(example[i]); j++ {
				f, _ := strconv.ParseFloat(example[i][j], 32)
				aux = append(aux, f)

			}
		}
		nNode = append(nNode, aux)
	}

	return
}

func firstEcode(nList []Node, nClusters int, maxIter int) string {

	sLong := toString(nList)

	t := toInt(nClusters)
	p := toInt(maxIter)

	sLong = sLong + ";" + t + ";" + p

	return sLong
}
func trainingEncode(nList []Node, nCentroids []Node, maxIter int) string {
	sLong := toString(nList)
	sCentroid := toString(nCentroids)
	p := toInt(maxIter)

	sLong = sLong + ";" + sCentroid + ";" + p

	return sLong
}

func firstDecode(sLong string) (recon [][]float64, newNCluster int, newNIter int) {

	deConvert := strings.Split(sLong, ";")

	deNode := strings.Split(deConvert[0], ",")

	fmt.Println(deNode[0])

	recon = decodeNode(deNode)

	newNCluster, _ = strconv.Atoi(deConvert[1])
	newNIter, _ = strconv.Atoi(deConvert[2])

	// var example [][]string
	// for i := 0; i < len(deNode); i++ {
	// 	s := strings.Split(deNode[i], "-")
	// 	example = append(example, s)
	// }

	// for i := 0; i < len(example)-1; i++ {
	// 	var aux []float64
	// 	if len(example[i]) != 0 {
	// 		for j := 0; j < len(example[i]); j++ {
	// 			f, _ := strconv.ParseFloat(example[i][j], 32)
	// 			aux = append(aux, f)

	// 		}
	// 	}
	// 	recon = append(recon, aux)
	// }

	return
}

func trainingDecode(sLong string) (recon [][]float64, newNCluster [][]float64, newNIter int) {

	deConvert := strings.Split(sLong, ";")

	deNode := strings.Split(deConvert[0], ",")

	deCluster := strings.Split(deConvert[1], ",")

	fmt.Println(deNode[1])

	// var example [][]string
	// for i := 0; i < len(deNode); i++ {
	// 	s := strings.Split(deNode[i], "-")
	// 	example = append(example, s)
	// }

	// for i := 0; i < len(example)-1; i++ {
	// 	var aux []float64
	// 	if len(example[i]) != 0 {
	// 		for j := 0; j < len(example[i]); j++ {
	// 			f, _ := strconv.ParseFloat(example[i][j], 32)
	// 			aux = append(aux, f)

	// 		}
	// 	}
	// 	recon = append(recon, aux)
	// }
	recon = decodeNode(deNode)
	newNCluster = decodeNode(deCluster)
	newNIter, _ = strconv.Atoi(deConvert[2])

	return
}

func KMeansInit(nList []Node, nClusters int, maxIter int) (bool, []Node) {
	lnList := len(nList)
	if lnList < nClusters {
		return false, nil
	}
	cons := 0
	for i, Node := range nList {
		nDims := len(Node)
		if i > 0 && len(Node) != cons {
			return false, nil
		}
		cons = nDims
	}
	nCentroids := make([]Node, nClusters)
	randN := rand.New(rand.NewSource(time.Now().UnixNano()))

	//Pick random centroids
	for i := 0; i < nClusters; i++ {
		srcIndex := randN.Intn(lnList)
		srcLen := len(nList[srcIndex])
		nCentroids[i] = make(Node, srcLen)
		copy(nCentroids[i], nList[randN.Intn(lnList)])
	}

	//Aquí iria la primera codificación s tring de los nodos
	sLong := firstEcode(nList, nClusters, maxIter)
	// la función KmeansTraining llamaría a sólo un String
	a, b, c := firstDecode(sLong)

	fmt.Println("Original array")
	fmt.Println(nList)
	fmt.Println("Original number of clusters")
	fmt.Println(nClusters)
	fmt.Println("Original number of iteration")
	fmt.Println(maxIter)
	fmt.Println("array coded")
	fmt.Println(sLong)
	fmt.Println("Coded/Decoded array")
	fmt.Println(a)
	fmt.Println("Coded/Decoded number of clusters")
	fmt.Println(b)
	fmt.Println("Coded/Decoded number of iteration")
	fmt.Println(c)

	//return KMeansTraining(nodeStrings)
	return KMeansTraining(nList, nCentroids, maxIter)
}

func KMeansTraining(nList []Node, nCentroids []Node, maxIter int) (bool, []Node) {
	//Aquí se clerarían las variables usadas sólo en la función.
	//Continuación iria la decodificación y divisón del string para asignarlo a sus variables

	sTraining := trainingEncode(nList, nCentroids, maxIter)
	fmt.Println(sTraining)
	_, b, _ := trainingDecode(sTraining)
	fmt.Println("Centroides decodeados")
	fmt.Println(b)

	//Training arc
	canMove := true
	for i := 0; i < maxIter && canMove; i++ {
		canMove = false
		cluster := make(map[int][]Node)
		for _, Node := range nList {
			nearest := Nearest(Node, nCentroids)
			cluster[nearest] = append(cluster[nearest], Node)
		}
		for key, value := range cluster {
			nNode := meanNode(value)
			if equals(nCentroids[key], nNode) == false {
				nCentroids[key] = nNode
				canMove = true
			}
		}
	}
	//codificar again en String los nodos
	return true, nCentroids
	//return false, KMeansTraining(nodeStrings)
}
