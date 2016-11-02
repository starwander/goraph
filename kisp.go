// Copyright(c) 2016 Ethan Zhuang <zhuangwj@gmail.com>.

package goraph

import (
	"math"
)

// top k independent shortest path
func (graph *Graph) Kisp(source, destination Id, topK int) ([]float64, [][]Id, error) {
	var err error
	var i, j, k int
	var dijkstraDist map[Id]float64
	var dijkstraPrev map[Id]Id
	distTopK := make([]float64, topK)
	pathTopK := make([][]Id, topK)
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
		for i = 0; i < k; i++ {
			for j = 0; j < len(pathTopK[i])-1; j++ {
				graph.DisableEdge(pathTopK[i][j], pathTopK[i][j+1])
			}
		}
		dijkstraDist, dijkstraPrev, _ = graph.Dijkstra(source)
		distTopK[k] = dijkstraDist[destination]
		pathTopK[k] = getPath(dijkstraPrev, destination)

		graph.Reset()
	}

	return distTopK, pathTopK, nil
}
