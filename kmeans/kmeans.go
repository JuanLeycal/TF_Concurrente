package kmeans

import (
	"fmt"
	"math/rand"
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

	for i := 0; i < len(nList); i++ {
		var aux []string
		for j := 0; j < len(nList[i]); j++ {
			s := fmt.Sprintf("%f", nList[i][j])
			// s = s + "basura"
			// fmt.Println(s)
			aux = append(aux, s)
		}
		sNode = append(sNode, aux)
		// fmt.Println(sNode[i])
	}

	//fmt.Println(sNode)

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

func main() {

}
