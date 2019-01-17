// Copyright(c) 2016 Ethan Zhuang <zhuangwj@gmail.com>.

package goraph

import (
	"math"
)

// Kisp gets top k shortest independent path between two vertex in the graph.
// Independent means no vertex is shared between path.
func (graph *Graph) Kisp(source, destination ID, topK int) ([]float64, [][]ID, error) {
	var err error
	var i, k int
	var dijkstraDist map[ID]float64
	var dijkstraPrev map[ID]ID
	distTopK := make([]float64, topK)
	pathTopK := make([][]ID, topK)
	for i := 0; i < topK; i++ {
		distTopK[i] = math.Inf(1)
	}

	dijkstraDist, dijkstraPrev, err = graph.Dijkstra(source)
	if err != nil {
		return nil, nil, err
	}
	distTopK[0] = dijkstraDist[destination]
	pathTopK[0] = getPath(dijkstraPrev, destination)

	for k = 1; k < topK && distTopK[k-1] != math.Inf(1); k++ {
		for i = 0; i < len(pathTopK[k-1])-1; i++ {
			graph.DisableEdge(pathTopK[k-1][i], pathTopK[k-1][i+1])
		}
		dijkstraDist, dijkstraPrev, _ = graph.Dijkstra(source)
		distTopK[k] = dijkstraDist[destination]
		pathTopK[k] = getPath(dijkstraPrev, destination)
	}
	graph.Reset()

	return distTopK, pathTopK, nil
}
