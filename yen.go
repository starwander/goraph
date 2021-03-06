// Copyright(c) 2016 Ethan Zhuang <zhuangwj@gmail.com>.

package goraph

import (
	"math"
	"sort"
)

type potential struct {
	dist float64
	path []ID
}

// Yen gets top k shortest loopless path between two vertex in the graph.
// https://en.wikipedia.org/wiki/Yen%27s_algorithm
func (graph *Graph) Yen(source, destination ID, topK int) ([]float64, [][]ID, error) {
	var err error
	var i, j, k int
	var dijkstraDist map[ID]float64
	var dijkstraPrev map[ID]ID
	var existed bool
	var spurWeight float64
	var spurPath []ID
	var potentials []potential
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

	for k = 1; k < topK; {
		for i = 0; i < len(pathTopK[k-1])-1; i++ {
			for j = 0; j < k; j++ {
				if isShareRootPath(pathTopK[j], pathTopK[k-1][:i+1]) {
					graph.DisableEdge(pathTopK[j][i], pathTopK[j][i+1])
				}
			}
			graph.DisablePath(pathTopK[k-1][:i])

			dijkstraDist, dijkstraPrev, _ = graph.Dijkstra(pathTopK[k-1][i])
			if dijkstraDist[destination] != math.Inf(1) {
				spurWeight = graph.GetPathWeight(pathTopK[k-1][:i+1]) + dijkstraDist[destination]
				spurPath = mergePath(pathTopK[k-1][:i], getPath(dijkstraPrev, destination))
				existed = false
				for _, each := range potentials {
					if isSamePath(each.path, spurPath) {
						existed = true
						break
					}
				}
				if !existed {
					potentials = append(potentials, potential{
						spurWeight,
						spurPath,
					})
				}
			}

			graph.Reset()
		}

		if len(potentials) == 0 {
			break
		}
		sort.Slice(potentials, func(i, j int) bool {
			return potentials[i].dist < potentials[j].dist
		})

		if len(potentials) >= topK-k {
			for l := 0; k < topK; l++ {
				distTopK[k] = potentials[l].dist
				pathTopK[k] = potentials[l].path
				k++
			}
			break
		} else {
			distTopK[k] = potentials[0].dist
			pathTopK[k] = potentials[0].path
			potentials = potentials[1:]
			k++
		}
	}

	return distTopK, pathTopK, nil
}

func isShareRootPath(path, rootPath []ID) bool {
	if len(path) < len(rootPath) {
		return false
	}

	return isSamePath(path[:len(rootPath)], rootPath)
}

func isSamePath(path1, path2 []ID) bool {
	if len(path1) != len(path2) {
		return false
	}

	for i := 0; i < len(path1); i++ {
		if path1[i] != path2[i] {
			return false
		}
	}

	return true
}

func mergePath(path1, path2 []ID) []ID {
	newPath := []ID{}
	newPath = append(newPath, path1...)
	newPath = append(newPath, path2...)

	return newPath
}
