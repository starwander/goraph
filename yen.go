// Copyright(c) 2016 Ethan Zhuang <zhuangwj@gmail.com>.

package goraph

import (
	"math"
)

// https://en.wikipedia.org/wiki/Yen%27s_algorithm
func (graph *Graph) Yen(source, destination Id, topK int) ([]float64, [][]Id, error) {
	var err error
	var i, j, k int
	var spurWeight float64
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

	for k = 1; k < topK; k++ {
		for i = 0; i < len(pathTopK[k-1])-1; i++ {
			for j = 0; j < k; j++ {
				if isShareRootPath(pathTopK[j], pathTopK[k-1][:i+1]) {
					graph.DisableEdge(pathTopK[j][i], pathTopK[j][i+1])
				}
			}
			graph.DisablePath(pathTopK[k-1][:i])

			dijkstraDist, dijkstraPrev, _ = graph.Dijkstra(pathTopK[k-1][i])
			spurWeight = graph.GetPathWeight(pathTopK[k-1][:i+1]) + dijkstraDist[destination]
			if spurWeight < distTopK[k] {
				distTopK[k] = spurWeight
				pathTopK[k] = mergePath(pathTopK[k-1][:i], getPath(dijkstraPrev, destination))
			}

			graph.Reset()
		}
	}

	return distTopK, pathTopK, nil
}

func isShareRootPath(path, rootPath []Id) bool {
	if len(path) < len(rootPath) {
		return false
	}

	for i := 0; i < len(rootPath); i++ {
		if path[i] != rootPath[i] {
			return false
		}
	}

	return true
}

func mergePath(path1, path2 []Id) []Id {
	newPath := []Id{}
	newPath = append(newPath, path1...)
	newPath = append(newPath, path2...)

	return newPath
}
