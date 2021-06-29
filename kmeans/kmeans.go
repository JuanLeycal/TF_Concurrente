package kmeans

import (
	"math/rand"
	"time"
)

func KMeans(nList []Node, nClusters int, maxIter int) (bool, []Node) {
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
	return true, nCentroids
}
