// Copyright(c) 2016 Ethan Zhuang <zhuangwj@gmail.com>.

package goraph

import (
	"fmt"
	"github.com/starwander/GoFibonacciHeap"
	"math"
)

// Dijkstra gets the shortest path from one vertex to all other vertices in the graph.
// https://en.wikipedia.org/wiki/Dijkstra%27s_algorithm
func (graph *Graph) Dijkstra(source Id) (dist map[Id]float64, prev map[Id]Id, err error) {
	if _, exists := graph.vertices[source]; !exists {
		return nil, nil, fmt.Errorf("Vertex %v is not existed", source)
	}

	dist = make(map[Id]float64)
	prev = make(map[Id]Id)
	heap := fibHeap.NewFibHeap()

	for id := range graph.vertices {
		prev[id] = nil
		if id != source {
			dist[id] = math.Inf(1)
			heap.Insert(id, math.Inf(1))
		} else {
			dist[id] = 0
			heap.Insert(id, 0)
		}
	}

	for heap.Num() != 0 {
		min, _ := heap.ExtractMin()
		for to, edge := range graph.egress[min] {
			if edge.getWeight() < 0 {
				return nil, nil, fmt.Errorf("Negative weight form vertex %v to vertex %v is not allowed", min, to)
			}
			if !edge.enable {
				continue
			}
			if dist[min]+edge.getWeight() < dist[to] {
				heap.DecreaseKey(to, dist[min]+edge.getWeight())
				prev[to] = min
				dist[to] = dist[min] + edge.getWeight()
			}
		}
	}

	return
}

func getPath(prev map[Id]Id, lastNode Id) (path []Id) {
	prevNode := prev[lastNode]
	if prevNode == nil {
		return nil
	}

	reversePath := []Id{lastNode}
	for ; prevNode != nil; prevNode = prev[prevNode] {
		reversePath = append(reversePath, prevNode)
	}

	path = make([]Id, len(reversePath))
	for index, node := range reversePath {
		path[len(reversePath)-index-1] = node
	}

	return
}
