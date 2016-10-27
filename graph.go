// Copyright(c) 2016 Ethan Zhuang <zhuangwj@gmail.com>.

package goraph

import (
	"fmt"
	"math"
)

type Id interface{}

type Vertex interface {
	Id() Id
	Out() map[Id]float64
	In() map[Id]float64
}

type Edge interface {
	Get() (from Id, to Id, weight float64)
}

type Graph struct {
	vertices map[Id]*vertex
	egress   map[Id]map[Id]*edge
	ingress  map[Id]map[Id]*edge
}

type vertex struct {
	this   interface{}
	enable bool
}

type edge struct {
	weight  float64
	enable  bool
	changed bool
}

func NewGraph() *Graph {
	graph := new(Graph)
	graph.vertices = make(map[Id]*vertex)
	graph.egress = make(map[Id]map[Id]*edge)
	graph.ingress = make(map[Id]map[Id]*edge)

	return graph
}

func (graph *Graph) GetVertex(id Id) (vertex interface{}, err error) {
	if v, exists := graph.vertices[id]; exists {
		vertex = v.this
		return
	}

	err = fmt.Errorf("Vertex %v is not found", id)
	return
}

func (graph *Graph) GetEdge(from Id, to Id) (float64, error) {
	if _, exists := graph.vertices[from]; !exists {
		return math.Inf(1), fmt.Errorf("Vertex(from) %v is not found", from)
	}

	if _, exists := graph.vertices[to]; !exists {
		return math.Inf(1), fmt.Errorf("Vertex(to) %v is not found", to)
	}

	if edge, exists := graph.egress[from][to]; exists {
		return edge.weight, nil
	}

	return math.Inf(1), nil
}

func (graph *Graph) AddVertex(id Id, v interface{}) error {
	if _, exists := graph.vertices[id]; exists {
		return fmt.Errorf("Vertex %v is duplicate", id)
	}

	graph.vertices[id] = &vertex{v, true}
	graph.egress[id] = make(map[Id]*edge)
	graph.ingress[id] = make(map[Id]*edge)

	return nil
}

func (graph *Graph) AddEdge(from Id, to Id, weight float64) error {
	if weight == math.Inf(-1) {
		return fmt.Errorf("-inf weight is reserved for internal usage")
	}

	if _, exists := graph.vertices[from]; !exists {
		return fmt.Errorf("Vertex(from) %v is not found", from)
	}

	if _, exists := graph.vertices[to]; !exists {
		return fmt.Errorf("Vertex(to) %v is not found", to)
	}

	if _, exists := graph.egress[from][to]; exists {
		return fmt.Errorf("Edge from %v to %v is duplicate", from, to)
	}

	graph.egress[from][to] = &edge{weight, true, false}
	graph.ingress[to][from] = graph.egress[from][to]

	return nil
}

func (graph *Graph) UpdateEdge(from Id, to Id, weight float64) error {
	if weight == math.Inf(-1) {
		return fmt.Errorf("-inf weight is reserved for internal usage")
	}

	if _, exists := graph.vertices[from]; !exists {
		return fmt.Errorf("Vertex(from) %v is not found", from)
	}

	if _, exists := graph.vertices[to]; !exists {
		return fmt.Errorf("Vertex(to) %v is not found", to)
	}

	if edge, exists := graph.egress[from][to]; exists {
		edge.weight = weight
		return nil
	}

	return fmt.Errorf("Edge from %v to %v is not found", from, to)
}

func (graph *Graph) DeleteVertex(id Id) interface{} {
	if vertex, exists := graph.vertices[id]; exists {
		for to := range graph.egress[id] {
			delete(graph.ingress[to], id)
		}
		for from := range graph.ingress[id] {
			delete(graph.egress[from], id)
		}
		delete(graph.egress, id)
		delete(graph.ingress, id)
		delete(graph.vertices, id)

		return vertex.this
	}

	return nil
}

func (graph *Graph) DeleteEdge(from Id, to Id) {
	if _, exists := graph.vertices[from]; !exists {
		return
	}

	if _, exists := graph.vertices[to]; !exists {
		return
	}

	if _, exists := graph.egress[from][to]; exists {
		delete(graph.egress[from], to)
		delete(graph.ingress[to], from)
	}
}

func (graph *Graph) AddVertexWithEdges(v Vertex) error {
	if _, exists := graph.vertices[v.Id()]; exists {
		return fmt.Errorf("Vertex %v is duplicate", v.Id())
	}

	graph.vertices[v.Id()] = &vertex{v, true}
	graph.egress[v.Id()] = make(map[Id]*edge)
	graph.ingress[v.Id()] = make(map[Id]*edge)

	for outTo, weight := range v.Out() {
		if weight == math.Inf(-1) {
			return fmt.Errorf("-inf weight is reserved for internal usage")
		}

		graph.egress[v.Id()][outTo] = &edge{weight, true, false}
		if _, exists := graph.ingress[outTo]; !exists {
			graph.ingress[outTo] = make(map[Id]*edge)
		}
		graph.ingress[outTo][v.Id()] = graph.egress[v.Id()][outTo]
	}

	for inFrom, weight := range v.In() {
		if weight == math.Inf(-1) {
			return fmt.Errorf("-inf weight is reserved for internal usage")
		}

		graph.ingress[v.Id()][inFrom] = &edge{weight, true, false}
		if _, exists := graph.egress[inFrom]; !exists {
			graph.egress[inFrom] = make(map[Id]*edge)
		}
		graph.egress[inFrom][v.Id()] = graph.ingress[v.Id()][inFrom]
	}

	return nil
}

func (graph *Graph) CheckIntegrity() error {
	for from, out := range graph.egress {
		if _, exists := graph.vertices[from]; !exists {
			return fmt.Errorf("Vertex %v is not found", from)
		}
		for to := range out {
			if _, exists := graph.vertices[to]; !exists {
				return fmt.Errorf("Vertex %v is not found", to)
			}
		}
	}

	for to, in := range graph.ingress {
		if _, exists := graph.vertices[to]; !exists {
			return fmt.Errorf("Vertex %v is not found", to)
		}
		for from := range in {
			if _, exists := graph.vertices[from]; !exists {
				return fmt.Errorf("Vertex %v is not found", from)
			}
		}
	}

	return nil
}

func (graph *Graph) GetPathWeight(path []Id) (totalWeight float64) {
	if len(path) < 2 {
		return math.Inf(-1)
	}

	for i := 0; i < len(path)-1; i++ {
		if _, exists := graph.vertices[path[i]]; !exists {
			return math.Inf(-1)
		}
		if edge, exists := graph.egress[path[i]][path[i+1]]; exists {
			totalWeight += edge.weight
		} else {
			return math.Inf(1)
		}
	}

	return totalWeight
}
