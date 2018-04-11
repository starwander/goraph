// Copyright(c) 2016 Ethan Zhuang <zhuangwj@gmail.com>.

package goraph

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"math"
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
			graph.AddEdge("F", "E", -1, nil)
			Expect(graph.CheckIntegrity()).ShouldNot(HaveOccurred())

			dist, path, err := graph.Yen("C", "H", 3)
			Expect(err).Should(HaveOccurred())
			Expect(dist).Should(BeNil())
			Expect(path).Should(BeNil())
		})

		It("Given a graph without negative edge, when call yen api, then get the top k shortest paths from the source vertex to the destination vertex in the graph.", func() {
			dist, path, err := graph.Yen("C", "H", 9)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(dist[0]).Should(BeEquivalentTo(5))
			Expect(path[0]).Should(BeEquivalentTo([]Id{"C", "E", "F", "H"}))
			Expect(dist[1]).Should(BeEquivalentTo(7))
			Expect(path[1]).Should(BeEquivalentTo([]Id{"C", "E", "G", "H"}))
			Expect(dist[2]).Should(BeEquivalentTo(8))
			Expect(path[2]).Should(BeEquivalentTo([]Id{"C", "D", "F", "H"}))
			Expect(dist[3]).Should(BeEquivalentTo(8))
			Expect(path[3]).Should(BeEquivalentTo([]Id{"C", "E", "F", "G", "H"}))
			Expect(dist[4]).Should(BeEquivalentTo(8))
			Expect(path[4]).Should(BeEquivalentTo([]Id{"C", "E", "D", "F", "H"}))
			Expect(dist[5]).Should(BeEquivalentTo(11))
			Expect(path[5]).Should(BeEquivalentTo([]Id{"C", "D", "F", "G", "H"}))
			Expect(dist[6]).Should(BeEquivalentTo(11))
			Expect(path[6]).Should(BeEquivalentTo([]Id{"C", "E", "D", "F", "G", "H"}))
			Expect(dist[7]).Should(BeEquivalentTo(math.Inf(1)))
			Expect(path[7]).Should(BeNil())
			Expect(dist[8]).Should(BeEquivalentTo(math.Inf(1)))
			Expect(path[8]).Should(BeNil())
		})

		It("Given another graph without negative edge, when call yen api, then get the top k shortest paths from the source vertex to the destination vertex in the graph.", func() {
			mygraph := NewGraph()
			mygraph.AddVertex("A", nil)
			mygraph.AddVertex("B", nil)
			mygraph.AddVertex("C", nil)
			mygraph.AddVertex("D", nil)
			mygraph.AddVertex("E", nil)
			mygraph.AddEdge("A", "B", 1, nil)
			mygraph.AddEdge("B", "C", 1, nil)
			mygraph.AddEdge("C", "D", 1, nil)
			mygraph.AddEdge("A", "D", 2, nil)
			mygraph.AddEdge("B", "E", 2, nil)
			mygraph.AddEdge("E", "D", 1, nil)
			dist, path, err := mygraph.Yen("A", "D", 3)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(dist[0]).Should(BeEquivalentTo(2))
			Expect(path[0]).Should(BeEquivalentTo([]Id{"A", "D"}))
			Expect(dist[1]).Should(BeEquivalentTo(3))
			Expect(path[1]).Should(BeEquivalentTo([]Id{"A", "B", "C", "D"}))
			Expect(dist[2]).Should(BeEquivalentTo(4))
			Expect(path[2]).Should(BeEquivalentTo([]Id{"A", "B", "E", "D"}))
		})

		It("Support early stop when there are enough potential paths in each iteration", func() {
			mygraph := NewGraph()
			mygraph.AddVertex("0", nil)
			mygraph.AddVertex("1", nil)
			mygraph.AddVertex("2", nil)
			mygraph.AddVertex("3", nil)
			mygraph.AddVertex("4", nil)
			mygraph.AddEdge("0", "1", 1, nil)
			mygraph.AddEdge("0", "2", 1, nil)
			mygraph.AddEdge("1", "2", 1, nil)
			mygraph.AddEdge("2", "3", 1, nil)
			mygraph.AddEdge("2", "4", 1, nil)
			mygraph.AddEdge("3", "4", 1, nil)
			dist, path, err := mygraph.Yen("0", "4", 3)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(dist[0]).Should(BeEquivalentTo(2))
			Expect(path[0]).Should(BeEquivalentTo([]Id{"0", "2", "4"}))
			Expect(dist[1]).Should(BeEquivalentTo(3))
			Expect(path[1]).Should(BeEquivalentTo([]Id{"0", "1", "2", "4"}))
			Expect(dist[2]).Should(BeEquivalentTo(3))
			Expect(path[2]).Should(BeEquivalentTo([]Id{"0", "2", "3", "4"}))
		})
	})

	Context("bugfix test", func() {
		BeforeEach(func() {
			graph = NewGraph()
		})

		AfterEach(func() {
			graph = nil
		})

		It("Issue reported by gsidebottom for missing path.", func() {
			graph.AddVertex("0", nil)
			graph.AddVertex("1", nil)
			graph.AddVertex("2", nil)
			graph.AddVertex("3", nil)
			graph.AddVertex("4", nil)
			graph.AddEdge("0", "1", 1, nil)
			graph.AddEdge("0", "2", 1, nil)
			graph.AddEdge("1", "2", 1, nil)
			graph.AddEdge("2", "3", 1, nil)
			graph.AddEdge("2", "4", 1, nil)
			graph.AddEdge("3", "4", 1, nil)
			dist, path, err := graph.Yen("0", "4", 5)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(dist[0]).Should(BeEquivalentTo(2))
			Expect(path[0]).Should(BeEquivalentTo([]Id{"0", "2", "4"}))
			Expect(dist[1]).Should(BeEquivalentTo(3))
			Expect(path[1]).Should(BeEquivalentTo([]Id{"0", "1", "2", "4"}))
			Expect(dist[2]).Should(BeEquivalentTo(3))
			Expect(path[2]).Should(BeEquivalentTo([]Id{"0", "2", "3", "4"}))
			Expect(dist[3]).Should(BeEquivalentTo(4))
			Expect(path[3]).Should(BeEquivalentTo([]Id{"0", "1", "2", "3", "4"}))
		})
	})
})
