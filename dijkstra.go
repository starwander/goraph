// Copyright(c) 2016 Ethan Zhuang <zhuangwj@gmail.com>.

package goraph

import (
	"fmt"
	"github.com/EthanZhuang/GoFibonacciHeap"
	"math"
)

// https://en.wikipedia.org/wiki/Dijkstra%27s_algorithm
//1  function Dijkstra(Graph, source):
//2      dist[source] ← 0                                    // Initialization
//3
//4      create vertex set Q
//5
//6      for each vertex v in Graph:
//7          if v ≠ source
//8              dist[v] ← INFINITY                          // Unknown distance from source to v
//9              prev[v] ← UNDEFINED                         // Predecessor of v
//10
//11         Q.add_with_priority(v, dist[v])
//12
//13
//14     while Q is not empty:                               // The main loop
//15         u ← Q.extract_min()                             // Remove and return best vertex
//16         for each neighbor v of u:                       // only v that is still in Q
//17             alt = dist[u] + length(u, v)
//18             if alt < dist[v]
//19                 dist[v] ← alt
//20                 prev[v] ← u
//21                 Q.decrease_priority(v, alt)
//22
//23     return dist[], prev[]
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
			if edge.weight < 0 {
				return nil, nil, fmt.Errorf("Negative weight form vertex %v to vertex %v is not allowed", min, to)
			}
			if !edge.enable {
				continue
			}
			if dist[min]+edge.weight < dist[to] {
				heap.DecreaseKey(to, dist[min]+edge.weight)
				prev[to] = min
				dist[to] = dist[min] + edge.weight
			}
		}
	}

	return
}
