package kmeans

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

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
	// la función KmeansTraining llamaría a sólo un String
	var sNode [][]string
	var pog []string

	for i := 0; i < len(nList); i++ {
		var aux []string
		for j := 0; j < len(nList[i]); j++ {
			s := fmt.Sprintf("%f", nList[i][j])
			// s = s + "basura"
			// fmt.Println(s)
			if j == len(nList[i])-1 {
				aux = append(aux, s+",")
			} else {
				aux = append(aux, s+"-")
			}

		}
		sNode = append(sNode, aux)
		// fmt.Println(sNode[i])
	}
	for i := 0; i < len(sNode); i++ {
		// for j := 0; j < len(sNode[i]); j++ {
		pog = append(pog, strings.Join(sNode[i], ""))
		// pog[i] = strings.Join(sNode[i], " ")
		//}
	}
	sLong := strings.Join(pog, "")
	//fmt.Println(sNode)
	//fmt.Println(pog)
	t := strconv.Itoa(nClusters)
	p := strconv.Itoa(maxIter)

	sLong = sLong + ";" + t + ";" + p

	//fmt.Println(sLong)
	deConvert := strings.Split(sLong, ";")
	//fmt.Println(deConvert)

	//fmt.Println(len(deConvert))
	deNode := strings.Split(deConvert[0], ",")
	fmt.Println(nList[0])
	fmt.Println(deNode[0])
	newNCluster, _ := strconv.Atoi(deConvert[1])
	newNIter, _ := strconv.Atoi(deConvert[2])

	var example [][]string
	for i := 0; i < len(deNode); i++ {
		s := strings.Split(deNode[i], "-")
		example = append(example, s)
	}
	//fmt.Println(example)

	var recon [][]float64

	for i := 0; i < len(example)-1; i++ {
		var aux []float64
		if len(example[i]) != 0 {
			for j := 0; j < len(example[i]); j++ {
				f, _ := strconv.ParseFloat(example[i][j], 32)
				aux = append(aux, f)

			}
		}
		recon = append(recon, aux)
		// fmt.Println(sNode[i])
	}
	fmt.Println("Original array")
	fmt.Println(nList)
	fmt.Println("Original number of clusters")
	fmt.Println(nClusters)
	fmt.Println("Original number of iteration")
	fmt.Println(maxIter)
	fmt.Println("Coded/Decoded array")
	fmt.Println(recon)
	fmt.Println("Coded/Decoded number of clusters")
	fmt.Println(newNCluster)
	fmt.Println("Coded/Decoded number of iteration")
	fmt.Println(newNIter)

	//return KMeansTraining(nodeStrings)
	return KMeansTraining(nList, nCentroids, maxIter)
}

func KMeansTraining(nList []Node, nCentroids []Node, maxIter int) (bool, []Node) {
	//Aquí se clerarían las variables usadas sólo en la función.
	//Continuación iria la decodificación y divisón del string para asignarlo a sus variables

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
