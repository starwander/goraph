// Copyright(c) 2016 Ethan Zhuang <zhuangwj@gmail.com>.

package goraph

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"math"
)

var _ = Describe("Tests of Kisp", func() {
	var (
		graph *Graph
	)

	Context("algorithem test", func() {
		BeforeEach(func() {
			graph = NewGraph()
			graph.AddVertexWithEdges(&myVertex{"C", map[ID]float64{"D": 3, "E": 2}, map[ID]float64{}})
			graph.AddVertexWithEdges(&myVertex{"D", map[ID]float64{"F": 4}, map[ID]float64{"C": 3, "E": 1}})
			graph.AddVertexWithEdges(&myVertex{"E", map[ID]float64{"D": 1, "F": 2, "G": 3}, map[ID]float64{"C": 2}})
			graph.AddVertexWithEdges(&myVertex{"F", map[ID]float64{"G": 2, "H": 1}, map[ID]float64{"D": 4, "E": 2}})
			graph.AddVertexWithEdges(&myVertex{"G", map[ID]float64{"H": 2}, map[ID]float64{"E": 3, "F": 2}})
			graph.AddVertexWithEdges(&myVertex{"H", map[ID]float64{}, map[ID]float64{"F": 1, "G": 2}})
			Expect(graph.CheckIntegrity()).ShouldNot(HaveOccurred())
		})

		AfterEach(func() {
			graph = nil
		})

		It("Given a non-negative edge graph, when call kisp api with non-existed verten, then get error.", func() {
			dist, path, err := graph.Kisp("S", "T", 3)
			Expect(err).Should(HaveOccurred())
			Expect(dist).Should(BeNil())
			Expect(path).Should(BeNil())
		})

		It("Given a graph with negative edge, when call kisp api, then get error.", func() {
			graph.AddEdge("F", "E", -1, nil)
			Expect(graph.CheckIntegrity()).ShouldNot(HaveOccurred())

			dist, path, err := graph.Kisp("C", "H", 3)
			Expect(err).Should(HaveOccurred())
			Expect(dist).Should(BeNil())
			Expect(path).Should(BeNil())
		})

		It("Given a non-negative edge graph, when call kisp api, then get the top k shortest paths from the source vertex to the destination vertex in the graph.", func() {
			dist, path, err := graph.Kisp("C", "H", 4)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(dist[0]).Should(BeEquivalentTo(5))
			Expect(path[0]).Should(BeEquivalentTo([]ID{"C", "E", "F", "H"}))
			Expect(dist[1]).Should(BeEquivalentTo(11))
			Expect(path[1]).Should(BeEquivalentTo([]ID{"C", "D", "F", "G", "H"}))
			Expect(dist[2]).Should(BeEquivalentTo(math.Inf(1)))
			Expect(path[2]).Should(BeNil())
			Expect(dist[3]).Should(BeEquivalentTo(math.Inf(1)))
			Expect(path[3]).Should(BeNil())
		})
	})
})
