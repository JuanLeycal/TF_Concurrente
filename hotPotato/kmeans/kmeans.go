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
	return strconv.Itoa(value)
}

func decodeNode(array []string) (nNode []Node) {
	var example [][]string
	for i := 0; i < len(array); i++ {
		s := strings.Split(array[i], "-")
		example = append(example, s)
	}
	for i := 0; i < len(example)-1; i++ {
		auxN := Node{}
		a, _ := strconv.ParseFloat(example[i][0], 32)
		b, _ := strconv.ParseFloat(example[i][1], 32)
		c, _ := strconv.ParseFloat(example[i][2], 32)
		d, _ := strconv.ParseFloat(example[i][3], 32)
		e, _ := strconv.ParseFloat(example[i][4], 32)
		auxN = append(auxN, a, b, c, d, e)
		nNode = append(nNode, auxN)
	}
	return
}
func TrainingEncode(nList []Node, nCentroids []Node, maxIter int) string {
	sLong := toString(nList)
	sCentroid := toString(nCentroids)
	p := toInt(maxIter)
	sLong = sLong + ";" + sCentroid + ";" + p
	return sLong
}

func TrainingDecode(sLong string) (recon []Node, newNCluster []Node, newNIter int) {
	deConvert := strings.Split(sLong, ";")
	deNode := strings.Split(deConvert[0], ",")
	deCluster := strings.Split(deConvert[1], ",")
	//fmt.Println(deNode[1])
	recon = decodeNode(deNode)
	newNCluster = decodeNode(deCluster)
	newNIter, _ = strconv.Atoi(deConvert[2])

	return
}

func KMeansInit(nList []Node, nClusters int, maxIter int) (bool, []Node, []Node, int) {
	lnList := len(nList)
	if lnList < nClusters {
		return false, nil, nil, 0
	}
	cons := 0
	for i, Node := range nList {
		nDims := len(Node)
		if i > 0 && len(Node) != cons {
			return false, nil, nil, 0
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

	//reemplazar v por encoding y llamada
	//xreturn KMeansTraining(nList, nCentroids, maxIter)

	//encode, send, wait, decode, poblar, retornar true, nCentroids

	return true, nList, nCentroids, maxIter
}

func KMeansTraining(nList []Node, nCentroids []Node, maxIter int) ([]Node, []Node, int) {
	//Aquí se clerarían las variables usadas sólo en la función.
	//Continuación iria la decodificación y divisón del string para asignarlo a sus variables (externo)

	//Training arc
	canMove := true
	//modificar para que haga una sola pasada
	if canMove {
		canMove = false
		cluster := make(map[int][]Node)
		for _, Node := range nList {
			nearest := Nearest(Node, nCentroids)
			cluster[nearest] = append(cluster[nearest], Node)
		}
		for key, value := range cluster {
			nNode := meanNode(value)
			if !equals(nCentroids[key], nNode) {
				nCentroids[key] = nNode
				canMove = true
			}
		}
	}
	//codificar again en String los nodos (externo)
	// retornar los mismos parametros, nList, nCentroids, maxIter
	return nList, nCentroids, maxIter

}
