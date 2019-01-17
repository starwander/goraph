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
	id     ID
	outTo  map[ID]float64
	inFrom map[ID]float64
}

type myEdge struct {
	from   ID
	to     ID
	weight float64
}

func (vertex *myVertex) ID() ID {
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

func (edge *myEdge) Get() (ID, ID, float64) {
	return edge.from, edge.to, edge.weight
}
