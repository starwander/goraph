## Golang Graph
[![Build Status](https://travis-ci.org/EthanZhuang/goraph.svg?branch=master)](https://travis-ci.org/EthanZhuang/goraph)
[![codecov](https://codecov.io/gh/EthanZhuang/goraph/branch/master/graph/badge.svg)](https://codecov.io/gh/EthanZhuang/goraph)
[![Go Report Card](https://goreportcard.com/badge/github.com/EthanZhuang/goraph)](https://goreportcard.com/report/github.com/EthanZhuang/goraph)
[![GoDoc](https://godoc.org/github.com/EthanZhuang/goraph?status.svg)](https://godoc.org/github.com/EthanZhuang/goraph)
[![License](https://img.shields.io/badge/license-Apache%202.0-blue.svg)](https://www.apache.org/licenses/LICENSE-2.0)

Goraph is a golang package provides basic graph structures and algorithms.

Goraph is NOT concurrent safe.

Current implemented(&radic;) and planned(&times;) algorithms:

| Algorithm |   BFS   |   DFS   | TopologicalSort | Kruskal  |     Prim    |   Dijkstra  |       Yen      |     Kisp      | BellmanFord |  FloydWarshall   |    EdmondsKarp    |
| :-------: | :-----: | :-----: | :-------------: | :------: | :---------: | :---------: | :------------: | :-----------: | :---------: | :--------------: | :---------------: |
|  Complex  | O(V+E)  | O(V+E)  |      O(V+E)     | O(ElogE) | O(E+VlogV)¹ | O(E+VlogV)¹ | O(KV(E+VlogV)) | O(K(E+VlogV)) |    O(VE)    | O(V<sup>3</sup>) | O(VE<sup>2</sup>) |
|  Status   | &times; | &times; |     &times;     | &times;  |   &times;   |   &radic;   |    &radic;     |    &radic;    |   &times;   |     &times;      |      &times;      |
¹ With Fibonacci heap.


* BFS: breadth first search.

* DFS: depth first search.

* TopologicalSort: is a linear ordering of a directed graph's vertices such that for every directed edge uv from vertex u to vertex v, u comes before v in the ordering.

* Kruskal: is a minimum-spanning-tree algorithm which finds an edge of the least possible weight that connects any two trees in the forest.

* Prim: is a greedy algorithm that finds a minimum spanning tree for a weighted undirected graph

* Dijkstra: computes shortest paths from a single source vertex to all of the other vertices in a graph with non-negative edge cost.

* Yen: computes K-shortest loopless paths between two vertex in a graph with non-negative edge cost.

* Kisp: computes K-shortest independent paths between two vertex in a graph with non-negative edge cost.

* BellmanFord: computes shortest paths from a single source vertex to all of the other vertices in a weighted digraph with positive or negative edge weights.

* FloydWarshall: computes all-pairs shortest paths in a weighted graph with positive or negative edge weights (but with no negative cycles).

* EdmondsKarp: computes the maximum flow in a flow network(graph).

##Requirements
#####Download this package

    go get github.com/EthanZhuang/GoFibonacciHeap
    go get github.com/EthanZhuang/goraph

#####Implements Value interface of this package for all values going to be inserted by value interfaces
```go
type Vertex interface {
	Id() Id
	Edges() []Edge
}

type Edge interface {
	Get() (from Id, to Id, weight float64)
}
```
## Supported Operations

* Graph operations:
 - GetVertex: get a vertex by input id.
 - GetEdge: gets the edge between the two vertices by input ids.
 - GetEdgeWeight: gets the weight of the edge between the two vertices by input ids.
 - AddVertex: adds a new vertex into the graph.
 - AddEdge: adds a new edge between the vertices by the input ids.
 - UpdateEdgeWeight: updates the weight of the edge between vertices by the input ids.
 - DeleteVertex: deletes a vertex from the graph and gets the value of the vertex.
 - DeleteEdge: deletes the edge between the vertices by the input id from the graph and gets the value of edge.
 - AddVertexWithEdges: adds a vertex value which implements Vertex interface.
 - CheckIntegrity: checks if any edge connects to or from unknown vertex.
 - GetPathWeight: gets the total weight along the path by input ids.
 - DisableEdge: disables the edge for further calculation.
 - DisableVertex: disables the vertex for further calculation.
 - DisablePath: disables all the vertices in the path for further calculation.
 - Reset: enables all vertices and edges for further calculation.

* Algorithm operations:
 - Dijkstra: gets the shortest path from one vertex to all other vertices in the graph.
 - Yen: gets top k shortest loopless path between two vertex in the graph.
 - Kisp: gets top k shortest independent path between two vertex in the graph.

## Example

```go
package main

import (
	"github.com/EthanZhuang/GoFibonacciHeap"
	"github.com/EthanZhuang/goraph"
)

```

## Reference

[GoDoc](https://godoc.org/github.com/EthanZhuang/goraph)

## LICENSE

goraph source code is licensed under the [Apache Licence, Version 2.0](http://www.apache.org/licenses/LICENSE-2.0.html).
