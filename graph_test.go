// Copyright(c) 2016 Ethan Zhuang <zhuangwj@gmail.com>.

package goraph

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"math"
)

var _ = Describe("Tests of Graph structure", func() {
	var (
		graph *Graph
	)

	Context("add/get vertex methods with Vertex interface tests", func() {
		BeforeEach(func() {
			graph = NewGraph()
		})

		AfterEach(func() {
			graph = nil
		})

		It("Given an empty graph, when get an vertex, then get a nil and error", func() {
			vertex, err := graph.GetVertex("S")
			Expect(err).Should(HaveOccurred())
			Expect(vertex).Should(BeNil())
		})

		It("Given an empty graph, when add a vertex, then can get the vertex by its id later", func() {
			myVertex := &myVertex{"S", map[Id]float64{"A": 10, "B": 10}, map[Id]float64{}}
			err := graph.AddVertexWithEdges(myVertex)
			Expect(err).ShouldNot(HaveOccurred())
			vertex, err := graph.GetVertex("S")
			Expect(vertex).ShouldNot(BeNil())
			Expect(vertex).Should(BeEquivalentTo(myVertex))
		})

		It("Given a graph, when add vertex with same ID, then get error", func() {
			graph.AddVertexWithEdges(&myVertex{"S", map[Id]float64{"A": 10, "B": 10}, map[Id]float64{}})
			err := graph.AddVertexWithEdges(&myVertex{"S", map[Id]float64{}, map[Id]float64{"S": 10, "B": 5}})
			Expect(err).Should(HaveOccurred())
		})

		It("Given a graph, when add a vertex with -inf weight edge, then get error", func() {
			err := graph.AddVertexWithEdges(&myVertex{"S", map[Id]float64{"T": math.Inf(-1)}, map[Id]float64{}})
			Expect(err).Should(HaveOccurred())
			err = graph.AddVertexWithEdges(&myVertex{"T", map[Id]float64{}, map[Id]float64{"S": math.Inf(-1)}})
			Expect(err).Should(HaveOccurred())
		})
	})

	Context("generic add/get vertex methods tests", func() {
		BeforeEach(func() {
			graph = NewGraph()
		})

		AfterEach(func() {
			graph = nil
		})

		It("Given an empty graph, when get a vertex, then get a nil and error", func() {
			vertex, err := graph.GetVertex("S")
			Expect(err).Should(HaveOccurred())
			Expect(vertex).Should(BeNil())
		})

		It("Given an empty graph, when add a vertex, then cat get the vertex by its id later", func() {
			err := graph.AddVertex("S", "I am vertex S")
			Expect(err).ShouldNot(HaveOccurred())
			vertex, err := graph.GetVertex("S")
			Expect(err).ShouldNot(HaveOccurred())
			Expect(vertex).Should(BeEquivalentTo("I am vertex S"))
		})

		It("Given a graph with vertex S, when add vertex S again, then get error", func() {
			err := graph.AddVertex("S", "I am vertex S")
			Expect(err).ShouldNot(HaveOccurred())
			err = graph.AddVertex("S", "I am vertex S too")
			Expect(err).Should(HaveOccurred())
		})
	})

	Context("add/get/update edge methods tests", func() {
		BeforeEach(func() {
			graph = NewGraph()
		})

		AfterEach(func() {
			graph = nil
		})

		It("Given a graph without vertex S, when get the edge from S, then get +inf and an error", func() {
			graph.AddVertexWithEdges(&myVertex{"A", map[Id]float64{}, map[Id]float64{"T": 10}})
			graph.AddVertexWithEdges(&myVertex{"T", map[Id]float64{"T": 10}, map[Id]float64{}})
			edge, err := graph.GetEdge("S", "T")
			Expect(err).Should(HaveOccurred())
			Expect(edge).Should(BeEquivalentTo(math.Inf(1)))
		})

		It("Given a graph without vertex T, when get the edge to T, then get +inf and an error", func() {
			graph.AddVertexWithEdges(&myVertex{"S", map[Id]float64{"T": 10}, map[Id]float64{}})
			graph.AddVertexWithEdges(&myVertex{"B", map[Id]float64{}, map[Id]float64{"T": 10}})
			edge, err := graph.GetEdge("S", "T")
			Expect(err).Should(HaveOccurred())
			Expect(edge).Should(BeEquivalentTo(math.Inf(1)))
		})

		It("Given a graph with S and T disconnected, when get the edge from S to T, then get +inf without error", func() {
			graph.AddVertexWithEdges(&myVertex{"S", map[Id]float64{"A": 10}, map[Id]float64{"A": 10}})
			graph.AddVertexWithEdges(&myVertex{"A", map[Id]float64{"S": 10}, map[Id]float64{"S": 10}})
			graph.AddVertexWithEdges(&myVertex{"B", map[Id]float64{"T": 10}, map[Id]float64{"T": 10}})
			graph.AddVertexWithEdges(&myVertex{"T", map[Id]float64{"B": 10}, map[Id]float64{"B": 10}})
			edge, err := graph.GetEdge("S", "T")
			Expect(err).ShouldNot(HaveOccurred())
			Expect(edge).Should(BeEquivalentTo(math.Inf(1)))
		})

		It("Given a graph with S and T connected, when get the edge from S to T, then get its weight without error", func() {
			graph.AddVertexWithEdges(&myVertex{"S", map[Id]float64{"T": 10}, map[Id]float64{"T": 10}})
			graph.AddVertexWithEdges(&myVertex{"T", map[Id]float64{"S": 10}, map[Id]float64{"S": 10}})
			edge, err := graph.GetEdge("S", "T")
			Expect(err).ShouldNot(HaveOccurred())
			Expect(edge).Should(BeEquivalentTo(10))
		})

		It("Given a graph with S and T connected, when get the edge from S to T, then get its weight without error", func() {
			graph.AddVertexWithEdges(&myVertex{"S", map[Id]float64{"T": 10}, map[Id]float64{"T": 10}})
			graph.AddVertexWithEdges(&myVertex{"T", map[Id]float64{"S": 10}, map[Id]float64{"S": 10}})
			edge, err := graph.GetEdge("S", "T")
			Expect(err).ShouldNot(HaveOccurred())
			Expect(edge).Should(BeEquivalentTo(10))
		})

		It("Given a graph without S, when add an edge from S, then get an error", func() {
			graph.AddVertex("T", "I am vertex T")
			err := graph.AddEdge("S", "T", 10)
			Expect(err).Should(HaveOccurred())
		})

		It("Given a graph without T, when add an edge to T, then get an error", func() {
			graph.AddVertex("S", "I am vertex S")
			err := graph.AddEdge("S", "T", 10)
			Expect(err).Should(HaveOccurred())
		})

		It("Given a graph with S and T disconnected, when add an edge from S to T with -inf weight, then get an error", func() {
			graph.AddVertex("S", "I am vertex S")
			graph.AddVertex("T", "I am vertex T")
			err := graph.AddEdge("S", "T", math.Inf(-1))
			Expect(err).Should(HaveOccurred())
		})

		It("Given a graph with S and T disconnected, when add an edge from S to T, then get nil error", func() {
			graph.AddVertex("S", "I am vertex S")
			graph.AddVertex("T", "I am vertex T")
			err := graph.AddEdge("S", "T", 10)
			Expect(err).ShouldNot(HaveOccurred())
			edge, err := graph.GetEdge("S", "T")
			Expect(err).ShouldNot(HaveOccurred())
			Expect(edge).Should(BeEquivalentTo(10))
		})

		It("Given a graph with S and T already connected, when add an edge from S to T again, then get an error", func() {
			graph.AddVertex("S", "I am vertex S")
			graph.AddVertex("T", "I am vertex T")
			err := graph.AddEdge("S", "T", 10)
			Expect(err).ShouldNot(HaveOccurred())
			err = graph.AddEdge("S", "T", 20)
			Expect(err).Should(HaveOccurred())
		})
	})

	Context("update methods tests", func() {
		BeforeEach(func() {
			graph = NewGraph()
		})

		AfterEach(func() {
			graph = nil
		})

		It("Given a graph without S, when update an edge from S, then get an error", func() {
			graph.AddVertex("T", "I am vertex T")
			err := graph.UpdateEdge("S", "T", 10)
			Expect(err).Should(HaveOccurred())
		})

		It("Given a graph without T, when update an edge to T, then get an error", func() {
			graph.AddVertex("S", "I am vertex S")
			err := graph.UpdateEdge("S", "T", 10)
			Expect(err).Should(HaveOccurred())
		})

		It("Given a graph with S and T disconnected, when update an edge from S to T, then get nil error", func() {
			graph.AddVertex("S", "I am vertex S")
			graph.AddVertex("T", "I am vertex T")
			err := graph.UpdateEdge("S", "T", 10)
			Expect(err).Should(HaveOccurred())
		})

		It("Given a graph with S and T connected, when add an edge from S to T with -inf weight, then get an error", func() {
			graph.AddVertex("S", "I am vertex S")
			graph.AddVertex("T", "I am vertex T")
			graph.AddEdge("S", "T", 10)
			err := graph.UpdateEdge("S", "T", math.Inf(-1))
			Expect(err).Should(HaveOccurred())
		})

		It("Given a graph with  S and T connected, when add an edge from S to T again, then get an error", func() {
			graph.AddVertex("S", "I am vertex S")
			graph.AddVertex("T", "I am vertex T")
			graph.AddEdge("S", "T", 10)
			err := graph.UpdateEdge("S", "T", 20)
			Expect(err).ShouldNot(HaveOccurred())
			edge, err := graph.GetEdge("S", "T")
			Expect(err).ShouldNot(HaveOccurred())
			Expect(edge).Should(BeEquivalentTo(20))
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
			graph.AddVertexWithEdges(&myVertex{"S", map[Id]float64{"A": 10, "B": 10}, map[Id]float64{}})
			err := graph.CheckIntegrity()
			Expect(err).Should(HaveOccurred())
			graph.AddVertexWithEdges(&myVertex{"A", map[Id]float64{}, map[Id]float64{"S": 10, "B": 5}})
			err = graph.CheckIntegrity()
			Expect(err).Should(HaveOccurred())
			graph.AddVertexWithEdges(&myVertex{"B", map[Id]float64{"A": 5}, map[Id]float64{"S": 10}})
			err = graph.CheckIntegrity()
			Expect(err).ShouldNot(HaveOccurred())
			graph.egress["C"] = map[Id]*edge{"B": {15, true, false}}
			err = graph.CheckIntegrity()
			Expect(err).Should(HaveOccurred())
			delete(graph.egress, "C")
			graph.ingress["S"]["T"] = &edge{20, true, false}
			err = graph.CheckIntegrity()
			Expect(err).Should(HaveOccurred())
			delete(graph.ingress["S"], "T")
			graph.ingress["T"] = map[Id]*edge{"S": {20, true, false}}
			err = graph.CheckIntegrity()
			Expect(err).Should(HaveOccurred())
		})
	})

	Context("delete methods tests", func() {
		BeforeEach(func() {
			graph = NewGraph()
			graph.AddVertexWithEdges(&myVertex{"S", map[Id]float64{"A": 10, "B": 10}, map[Id]float64{}})
			graph.AddVertexWithEdges(&myVertex{"A", map[Id]float64{}, map[Id]float64{"S": 10, "B": 5}})
			graph.AddVertexWithEdges(&myVertex{"B", map[Id]float64{"A": 5}, map[Id]float64{"S": 10}})
		})

		AfterEach(func() {
			graph = nil
		})

		It("Given a graph without vertex T, when delete vertex T, then get an nil", func() {
			v := graph.DeleteVertex("T")
			Expect(v).Should(BeNil())
		})

		It("Given a graph with vertex S, when delete vertex S, then get S and can not get S later", func() {
			v := graph.DeleteVertex("S")
			Expect(v).Should(BeEquivalentTo(&myVertex{"S", map[Id]float64{"A": 10, "B": 10}, map[Id]float64{}}))
			_, err := graph.GetVertex("S")
			Expect(err).Should(HaveOccurred())
			err = graph.CheckIntegrity()
			Expect(err).ShouldNot(HaveOccurred())
			v = graph.DeleteVertex("A")
			Expect(v).Should(BeEquivalentTo(&myVertex{"A", map[Id]float64{}, map[Id]float64{"S": 10, "B": 5}}))
			_, err = graph.GetVertex("A")
			Expect(err).Should(HaveOccurred())
			err = graph.CheckIntegrity()
			Expect(err).ShouldNot(HaveOccurred())
		})

		It("Given a graph without vertex T, when delete edge from/to vertex T, then nothing happens", func() {
			graph.DeleteEdge("T", "S")
			_, err := graph.GetVertex("S")
			Expect(err).ShouldNot(HaveOccurred())
			err = graph.CheckIntegrity()
			Expect(err).ShouldNot(HaveOccurred())
			_, err = graph.GetVertex("S")
			Expect(err).ShouldNot(HaveOccurred())
			graph.DeleteEdge("S", "T")
			_, err = graph.GetVertex("S")
			Expect(err).ShouldNot(HaveOccurred())
			err = graph.CheckIntegrity()
			Expect(err).ShouldNot(HaveOccurred())
		})

		It("Given a graph without vertex A and B connected, when delete edge between A and B, then the weight between A and B will be +inf", func() {
			weight, _ := graph.GetEdge("B", "A")
			Expect(weight).Should(BeEquivalentTo(5))
			graph.DeleteEdge("B", "A")
			weight, _ = graph.GetEdge("B", "A")
			Expect(weight).Should(BeEquivalentTo(math.Inf(1)))
			err := graph.CheckIntegrity()
			Expect(err).ShouldNot(HaveOccurred())
		})
	})

	Context("get total weight of path tests", func() {
		BeforeEach(func() {
			graph = NewGraph()
			graph.AddVertexWithEdges(&myVertex{"S", map[Id]float64{"A": 10, "B": 10}, map[Id]float64{}})
			graph.AddVertexWithEdges(&myVertex{"A", map[Id]float64{}, map[Id]float64{"S": 10, "B": 5}})
			graph.AddVertexWithEdges(&myVertex{"B", map[Id]float64{"A": 5}, map[Id]float64{"S": 10}})
		})

		AfterEach(func() {
			graph = nil
		})

		It("Given an empty path, when get its weight, then get -inf", func() {
			path := []Id{}
			Expect(graph.GetPathWeight(path)).Should(BeEquivalentTo(math.Inf(-1)))
		})

		It("Given a path contains unexisted vertex, when get its weight, then get -inf", func() {
			path := []Id{"T"}
			Expect(graph.GetPathWeight(path)).Should(BeEquivalentTo(math.Inf(-1)))
			path = []Id{"S", "A", "T"}
			Expect(graph.GetPathWeight(path)).Should(BeEquivalentTo(math.Inf(-1)))
		})

		It("Given a path only have one vertex, when get its weight, then get 0", func() {
			path := []Id{"S"}
			Expect(graph.GetPathWeight(path)).Should(BeEquivalentTo(0))
		})

		It("Given a path with vertex disconnected, when get its weight, then get +inf", func() {
			path := []Id{"B", "A", "S"}
			Expect(graph.GetPathWeight(path)).Should(BeEquivalentTo(math.Inf(1)))
		})

		It("Given a path with all vertex connected, when get its weight, then get the end to end weight of the path", func() {
			path := []Id{"S", "B", "A"}
			Expect(graph.GetPathWeight(path)).Should(BeEquivalentTo(15))
		})
	})
})
