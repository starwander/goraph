// Copyright(c) 2016 Ethan Zhuang <zhuangwj@gmail.com>.

package goraph

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Tests of Dijkstra", func() {
	var (
		graph *Graph
	)

	Context("exception test", func() {
		BeforeEach(func() {
			graph = NewGraph()
		})

		AfterEach(func() {
			graph = nil
		})

		It("Given a graph without vertex X, when call dijkstra api with X, then get two nil and error", func() {
			graph.AddVertexWithEdges(&myVertex{"S", map[Id]float64{"A": 10, "B": 10}, map[Id]float64{}})
			graph.AddVertexWithEdges(&myVertex{"A", map[Id]float64{}, map[Id]float64{"S": 10, "B": 5}})
			graph.AddVertexWithEdges(&myVertex{"B", map[Id]float64{"A": 5}, map[Id]float64{"S": 10}})

			dist, prev, err := graph.Dijkstra("X")
			Expect(dist).Should(BeNil())
			Expect(prev).Should(BeNil())
			Expect(err).Should(HaveOccurred())
		})

		It("Given a graph with negative edge, when call dijkstra api, then get two nil and error", func() {
			graph.AddVertexWithEdges(&myVertex{"S", map[Id]float64{"A": 10, "B": 10}, map[Id]float64{}})
			graph.AddVertexWithEdges(&myVertex{"A", map[Id]float64{}, map[Id]float64{"S": 10, "B": -5}})
			graph.AddVertexWithEdges(&myVertex{"B", map[Id]float64{"A": -5}, map[Id]float64{"S": 10}})

			dist, prev, err := graph.Dijkstra("S")
			Expect(dist).Should(BeNil())
			Expect(prev).Should(BeNil())
			Expect(err).Should(HaveOccurred())
		})
	})

	Context("algorithem test", func() {
		BeforeEach(func() {
			graph = NewGraph()
			graph.AddVertexWithEdges(&myVertex{"S", map[Id]float64{"B": 14}, map[Id]float64{"A": 15, "B": 14, "C": 9}})
			graph.AddVertexWithEdges(&myVertex{"A", map[Id]float64{"S": 15, "B": 5, "D": 20, "T": 44}, map[Id]float64{"B": 5, "D": 20, "T": 44}})
			graph.AddVertexWithEdges(&myVertex{"B", map[Id]float64{"S": 14, "A": 5, "D": 30, "E": 18}, map[Id]float64{"S": 14, "A": 5, "D": 30, "E": 18}})
			graph.AddVertexWithEdges(&myVertex{"C", map[Id]float64{"S": 9, "E": 24}, map[Id]float64{"E": 24}})
			graph.AddVertexWithEdges(&myVertex{"D", map[Id]float64{"A": 20, "B": 30, "E": 2, "F": 11, "T": 16}, map[Id]float64{"A": 20, "B": 30, "E": 2, "F": 11, "T": 16}})
			graph.AddVertexWithEdges(&myVertex{"E", map[Id]float64{"B": 18, "C": 24, "D": 2, "F": 6, "T": 19}, map[Id]float64{"B": 18, "C": 24, "D": 2, "F": 6, "T": 19}})
			graph.AddVertexWithEdges(&myVertex{"F", map[Id]float64{"D": 11, "E": 6, "T": 6}, map[Id]float64{"D": 11, "E": 6, "T": 6}})
			graph.AddVertexWithEdges(&myVertex{"T", map[Id]float64{"A": 44, "D": 16, "E": 19, "F": 6}, map[Id]float64{"A": 44, "D": 16, "E": 19, "F": 6}})
			Expect(graph.CheckIntegrity()).ShouldNot(HaveOccurred())
		})

		AfterEach(func() {
			graph = nil
		})

		It("Given a non-negative edge graph, when call dijkstra api with source vertex, then get the shortest paths from the source vertices to all other vertex in the graph.", func() {
			expectedDist := map[Id]float64{
				"S": 0,
				"B": 14,
				"A": 19,
				"E": 32,
				"D": 34,
				"F": 38,
				"T": 44,
				"C": 56,
			}
			expectedPrev := map[Id]Id{
				"S": nil,
				"B": "S",
				"A": "B",
				"E": "B",
				"D": "E",
				"F": "E",
				"T": "F",
				"C": "E",
			}

			Expect(graph.CheckIntegrity()).ShouldNot(HaveOccurred())
			dist, prev, err := graph.Dijkstra("S")
			Expect(err).ShouldNot(HaveOccurred())
			Expect(dist).Should(BeEquivalentTo(expectedDist))
			Expect(prev).Should(BeEquivalentTo(expectedPrev))
		})

		It("Given a graph with some edges disabled, when call dijkstra api with source vertex, then the disabled edges will not be calculated.", func() {
			graph.egress["A"]["B"].enable = false
			graph.egress["E"]["F"].enable = false
			expectedDist := map[Id]float64{
				"S": 0,
				"B": 14,
				"A": 19,
				"E": 32,
				"D": 34,
				"F": 45,
				"T": 50,
				"C": 56,
			}
			expectedPrev := map[Id]Id{
				"S": nil,
				"B": "S",
				"A": "B",
				"E": "B",
				"D": "E",
				"F": "D",
				"T": "D",
				"C": "E",
			}

			Expect(graph.CheckIntegrity()).ShouldNot(HaveOccurred())
			dist, prev, err := graph.Dijkstra("S")
			Expect(err).ShouldNot(HaveOccurred())
			Expect(dist).Should(BeEquivalentTo(expectedDist))
			Expect(prev).Should(BeEquivalentTo(expectedPrev))
		})
	})
})
