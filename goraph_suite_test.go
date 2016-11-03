// Copyright(c) 2016 Ethan Zhuang <zhuangwj@gmail.com>.

package goraph

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func TestProxy(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Goraph Suite")
}

var _ = Describe("Test initialization", func() {
	Context("Register suite setup and teardown function", func() {
		BeforeSuite(func() {
		})

		AfterSuite(func() {
		})
	})
})

type myVertex struct {
	id     Id
	outTo  map[Id]float64
	inFrom map[Id]float64
}

type myEdge struct {
	from   Id
	to     Id
	weight float64
}

func (vertex *myVertex) Id() Id {
	return vertex.id
}

func (vertex *myVertex) Edges() (edges []Edge) {
	edges = make([]Edge, len(vertex.outTo)+len(vertex.inFrom))
	i := 0
	for to, weight := range vertex.outTo {
		edges[i] = &myEdge{vertex.id, to, weight}
		i++
	}
	for from, weight := range vertex.inFrom {
		edges[i] = &myEdge{from, vertex.id, weight}
		i++
	}
	return
}

func (edge *myEdge) Get() (Id, Id, float64) {
	return edge.from, edge.to, edge.weight
}
