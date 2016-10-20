// Copyright(c) 2016 Ethan Zhuang <zhuangwj@gmail.com>.

package goraph

import (
	"fmt"
)

type Id interface{}

type Vertex interface {
	Id() Id
	Out() map[Id]float64
	In() map[Id]float64
}

type Edge interface {
	From() Id
	To() Id
	Weight() float64
}

type Graph struct {
	vertices map[Id]*vertex
	egress   map[Id]map[Id]*edge
	ingress  map[Id]map[Id]*edge
}

type vertex struct {
	this   Vertex
	enable bool
}

type edge struct {
	weight float64
	enable bool
}

func NewGraph() *Graph {
	graph := new(Graph)
	graph.vertices = make(map[Id]*vertex)
	graph.egress = make(map[Id]map[Id]*edge)
	graph.ingress = make(map[Id]map[Id]*edge)

	return graph
}

func (graph *Graph) GetVertex(id Id) (vertex Vertex, err error) {
	if v, exists := graph.vertices[id]; exists {
		vertex = v.this
		return
	}

	err = fmt.Errorf("Vertex %v is not found", id)
	return
}

func (graph *Graph) AddVertex(v Vertex) error {
	if _, exists := graph.vertices[v.Id()]; exists {
		return fmt.Errorf("Vertex %v is duplicate", v.Id())
	}

	graph.vertices[v.Id()] = &vertex{v, true}
	graph.egress[v.Id()] = make(map[Id]*edge)
	graph.ingress[v.Id()] = make(map[Id]*edge)

	for outTo, weight := range v.Out() {
		graph.egress[v.Id()][outTo] = &edge{weight, true}
		if _, exists := graph.ingress[outTo]; !exists {
			graph.ingress[outTo] = make(map[Id]*edge)
		}
		graph.ingress[outTo][v.Id()] = graph.egress[v.Id()][outTo]
	}

	for inFrom, weight := range v.In() {
		graph.ingress[v.Id()][inFrom] = &edge{weight, true}
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
