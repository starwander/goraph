// Copyright(c) 2016 Ethan Zhuang <zhuangwj@gmail.com>.

package goraph

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Tests of Yen", func() {
	var (
		graph *Graph
	)

	Context("algorithem test", func() {
		BeforeEach(func() {
			graph = NewGraph()
			graph.AddVertexWithEdges(&myVertex{"C", map[Id]float64{"D": 3, "E": 2}, map[Id]float64{}})
			graph.AddVertexWithEdges(&myVertex{"D", map[Id]float64{"F": 4}, map[Id]float64{"C": 3, "E": 1}})
			graph.AddVertexWithEdges(&myVertex{"E", map[Id]float64{"D": 1, "F": 2, "G": 3}, map[Id]float64{"C": 2}})
			graph.AddVertexWithEdges(&myVertex{"F", map[Id]float64{"G": 2, "H": 1}, map[Id]float64{"D": 4, "E": 2}})
			graph.AddVertexWithEdges(&myVertex{"G", map[Id]float64{"H": 2}, map[Id]float64{"E": 3, "F": 2}})
			graph.AddVertexWithEdges(&myVertex{"H", map[Id]float64{}, map[Id]float64{"F": 1, "G": 2}})
			Expect(graph.CheckIntegrity()).ShouldNot(HaveOccurred())
		})

		AfterEach(func() {
			graph = nil
		})

		It("Given a non-negative edge graph, when call yen api with non-existed verten, then get error.", func() {
			dist, path, err := graph.Yen("S", "T", 3)
			Expect(err).Should(HaveOccurred())
			Expect(dist).Should(BeNil())
			Expect(path).Should(BeNil())
		})

		It("Given a graph with negative edge, when call yen api, then get error.", func() {
			graph.AddEdge("F", "E", -1)
			Expect(graph.CheckIntegrity()).ShouldNot(HaveOccurred())

			dist, path, err := graph.Yen("C", "H", 3)
			Expect(err).Should(HaveOccurred())
			Expect(dist).Should(BeNil())
			Expect(path).Should(BeNil())
		})

		It("Given a non-negative edge graph, when call yen api, then get the top k shortest paths from the source vertex to the destination vertex in the graph.", func() {
			dist, path, err := graph.Yen("C", "H", 3)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(dist[0]).Should(BeEquivalentTo(5))
			Expect(path[0]).Should(BeEquivalentTo([]Id{"C", "E", "F", "H"}))
			Expect(dist[1]).Should(BeEquivalentTo(7))
			Expect(path[1]).Should(BeEquivalentTo([]Id{"C", "E", "G", "H"}))
			Expect(dist[2]).Should(BeEquivalentTo(8))
			Expect(path[2]).Should(BeEquivalentTo([]Id{"C", "D", "F", "H"}))
		})
	})
})
