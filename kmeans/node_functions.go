package kmeans

import "math"

type Node []float32

func meanNode(nList []Node) Node {
	mNode := make(Node, len(nList[0]))
	for _, v := range nList {
		for j := 0; j < len(mNode); j++ {
			mNode[j] += v[j]
		}
	}
	for i, v := range mNode {
		mNode[i] = v / float32(len(nList))
	}
	return mNode
}

func Nearest(n Node, nList []Node) int {
	count := len(nList)
	res := make(Node, count)
	ch := make(chan int)
	for i, node := range nList {
		go func(i int, node, cl Node) {
			res[i] = eucDistance(n, node)
			ch <- 1
		}(i, node, n)
	}
	wait(ch, res)
	nDist := res[0]
	minI := 0
	for i, dist := range res {
		if dist < nDist {
			nDist = dist
			minI = i
		}
	}
	return minI
}

func eucDistance(n1, n2 Node) float32 {
	ln1 := len(n1)
	square := make(Node, ln1, ln1)
	cont := make(chan int)
	for i := range n1 {
		go func(i int) {
			diff := n1[i] - n2[i]
			square[i] = float32(math.Pow(float64(diff), 2))
			cont <- 1
		}(i)
	}
	wait(cont, square)

	sum := float32(0)
	for _, v := range square {
		sum += v
	}
	return sum
}

func equals(n1, n2 Node) bool {
	ln1, ln2 := len(n1), len(n2)
	if ln1 != ln2 {
		return false
	}
	for i, v := range n1 {
		if v != n2[i] {
			return false
		}
	}
	return true
}

func wait(ch chan int, n Node) {
	count := len(n)
	<-ch
	for nDim := 1; nDim < count; nDim++ {
		<-ch
	}
}
