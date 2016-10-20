// Copyright(c) 2016 Ethan Zhuang <zhuangwj@gmail.com>.

package goraph

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Tests of Graph structure", func() {
	var (
		graph *Graph
	)

	Context("AddVertex/GetVertex tests", func() {
		BeforeEach(func() {
			graph = NewGraph()
		})

		AfterEach(func() {
			graph = nil
		})

		It("Given a empty graph, when get an vertex, then get a nil and error", func() {
			vertex, err := graph.GetVertex("S")
			Expect(err).Should(HaveOccurred())
			Expect(vertex).Should(BeNil())
		})

		It("Given a empty graph, when add an vertex, then cat get the vertex by its id later", func() {
			myVertex := &myVertex{"S", map[Id]float64{"A": 10, "B": 10}, map[Id]float64{}}
			err := graph.AddVertex(myVertex)
			Expect(err).ShouldNot(HaveOccurred())
			vertex, err := graph.GetVertex("S")
			Expect(vertex).ShouldNot(BeNil())
			Expect(vertex).Should(BeEquivalentTo(myVertex))
		})

		It("Given a graph, when add 2 vertex with same ID, then get error", func() {
			graph.AddVertex(&myVertex{"S", map[Id]float64{"A": 10, "B": 10}, map[Id]float64{}})
			err := graph.AddVertex(&myVertex{"S", map[Id]float64{}, map[Id]float64{"S": 10, "B": 5}})
			Expect(err).Should(HaveOccurred())
		})
	})

	Context("IntegrityCheck tests", func() {
		BeforeEach(func() {
			graph = NewGraph()
		})

		AfterEach(func() {
			graph = nil
		})

		It("Given a graph with an edge to an unknown vertex, when check integrity, then get error", func() {
			graph.AddVertex(&myVertex{"S", map[Id]float64{"A": 10, "B": 10}, map[Id]float64{}})
			err := graph.CheckIntegrity()
			Expect(err).Should(HaveOccurred())
			graph.AddVertex(&myVertex{"A", map[Id]float64{}, map[Id]float64{"S": 10, "B": 5}})
			err = graph.CheckIntegrity()
			Expect(err).Should(HaveOccurred())
			graph.AddVertex(&myVertex{"B", map[Id]float64{"A": 5}, map[Id]float64{"S": 10}})
			err = graph.CheckIntegrity()
			Expect(err).ShouldNot(HaveOccurred())
		})
	})
})
