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
			myVertex := &myVertex{"S", map[ID]float64{"A": 10, "B": 10}, map[ID]float64{}}
			err := graph.AddVertexWithEdges(myVertex)
			Expect(err).ShouldNot(HaveOccurred())
			vertex, err := graph.GetVertex("S")
			Expect(err).ShouldNot(HaveOccurred())
			Expect(vertex).ShouldNot(BeNil())
			Expect(vertex).Should(BeEquivalentTo(myVertex))
		})

		It("Given a graph, when add vertex with same ID, then get error", func() {
			graph.AddVertexWithEdges(&myVertex{"S", map[ID]float64{"A": 10, "B": 10}, map[ID]float64{}})
			err := graph.AddVertexWithEdges(&myVertex{"S", map[ID]float64{}, map[ID]float64{"S": 10, "B": 5}})
			Expect(err).Should(HaveOccurred())
		})

		It("Given a graph, when add a vertex with -inf weight edge, then get error", func() {
			err := graph.AddVertexWithEdges(&myVertex{"S", map[ID]float64{"T": math.Inf(-1)}, map[ID]float64{}})
			Expect(err).Should(HaveOccurred())
			err = graph.AddVertexWithEdges(&myVertex{"T", map[ID]float64{}, map[ID]float64{"S": math.Inf(-1)}})
			Expect(err).Should(HaveOccurred())
		})

		It("Given a graph, when add a vertex with unrelated edge, then get error", func() {
			err := graph.AddVertexWithEdges(&testVertex{"X", &myVertex{"S", map[ID]float64{"T": 10}, map[ID]float64{}}})
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

		It("Given a graph without vertex S, when get the edge(weight) from S, then get +inf and an error", func() {
			graph.AddVertexWithEdges(&myVertex{"A", map[ID]float64{}, map[ID]float64{"T": 10}})
			graph.AddVertexWithEdges(&myVertex{"T", map[ID]float64{"T": 10}, map[ID]float64{}})
			weight, err := graph.GetEdgeWeight("S", "T")
			Expect(err).Should(HaveOccurred())
			Expect(weight).Should(BeEquivalentTo(math.Inf(1)))
			edge, err := graph.GetEdge("S", "T")
			Expect(err).Should(HaveOccurred())
			Expect(edge).Should(BeNil())
		})

		It("Given a graph without vertex T, when get the edge(weight) to T, then get +inf and an error", func() {
			graph.AddVertexWithEdges(&myVertex{"S", map[ID]float64{"T": 10}, map[ID]float64{}})
			graph.AddVertexWithEdges(&myVertex{"B", map[ID]float64{}, map[ID]float64{"T": 10}})
			weight, err := graph.GetEdgeWeight("S", "T")
			Expect(err).Should(HaveOccurred())
			Expect(weight).Should(BeEquivalentTo(math.Inf(1)))
			edge, err := graph.GetEdge("S", "T")
			Expect(err).Should(HaveOccurred())
			Expect(edge).Should(BeNil())
		})

		It("Given a graph with S and T disconnected, when get the edge(weight) from S to T, then get +inf without error", func() {
			graph.AddVertexWithEdges(&myVertex{"S", map[ID]float64{"A": 10}, map[ID]float64{"A": 10}})
			graph.AddVertexWithEdges(&myVertex{"A", map[ID]float64{"S": 10}, map[ID]float64{"S": 10}})
			graph.AddVertexWithEdges(&myVertex{"B", map[ID]float64{"T": 10}, map[ID]float64{"T": 10}})
			graph.AddVertexWithEdges(&myVertex{"T", map[ID]float64{"B": 10}, map[ID]float64{"B": 10}})
			weight, err := graph.GetEdgeWeight("S", "T")
			Expect(err).ShouldNot(HaveOccurred())
			Expect(weight).Should(BeEquivalentTo(math.Inf(1)))
			edge, err := graph.GetEdge("S", "T")
			Expect(err).Should(HaveOccurred())
			Expect(edge).Should(BeNil())
		})

		It("Given a graph with S and T connected, when get the edge from S to T, then get its weight without error", func() {
			graph.AddVertexWithEdges(&myVertex{"S", map[ID]float64{"T": 10}, map[ID]float64{"T": 10}})
			graph.AddVertexWithEdges(&myVertex{"T", map[ID]float64{"S": 10}, map[ID]float64{"S": 10}})
			edge, err := graph.GetEdgeWeight("S", "T")
			Expect(err).ShouldNot(HaveOccurred())
			Expect(edge).Should(BeEquivalentTo(10))
		})

		It("Given a graph with S and T connected, when get the edge from S to T, then get its weight without error", func() {
			graph.AddVertexWithEdges(&myVertex{"S", map[ID]float64{"T": 10}, map[ID]float64{"T": 10}})
			graph.AddVertexWithEdges(&myVertex{"T", map[ID]float64{"S": 10}, map[ID]float64{"S": 10}})
			edge, err := graph.GetEdgeWeight("S", "T")
			Expect(err).ShouldNot(HaveOccurred())
			Expect(edge).Should(BeEquivalentTo(10))
		})

		It("Given a graph without S, when add an edge from S, then get an error", func() {
			graph.AddVertex("T", "I am vertex T")
			err := graph.AddEdge("S", "T", 10, nil)
			Expect(err).Should(HaveOccurred())
		})

		It("Given a graph without T, when add an edge to T, then get an error", func() {
			graph.AddVertex("S", "I am vertex S")
			err := graph.AddEdge("S", "T", 10, nil)
			Expect(err).Should(HaveOccurred())
		})

		It("Given a graph with S and T disconnected, when add an edge weight from S to T with -inf weight, then get an error", func() {
			graph.AddVertex("S", "I am vertex S")
			graph.AddVertex("T", "I am vertex T")
			err := graph.AddEdge("S", "T", math.Inf(-1), nil)
			Expect(err).Should(HaveOccurred())
		})

		It("Given a graph with S and T disconnected, when add an edge weight from S to T, then get nil error", func() {
			graph.AddVertex("S", "I am vertex S")
			graph.AddVertex("T", "I am vertex T")
			err := graph.AddEdge("S", "T", 10, nil)
			Expect(err).ShouldNot(HaveOccurred())
			edge, err := graph.GetEdgeWeight("S", "T")
			Expect(err).ShouldNot(HaveOccurred())
			Expect(edge).Should(BeEquivalentTo(10))
		})

		It("Given a graph with S and T already connected, when add an edge weight from S to T again, then get an error", func() {
			graph.AddVertex("S", "I am vertex S")
			graph.AddVertex("T", "I am vertex T")
			err := graph.AddEdge("S", "T", 10, nil)
			Expect(err).ShouldNot(HaveOccurred())
			err = graph.AddEdge("S", "T", 20, nil)
			Expect(err).Should(HaveOccurred())
		})

		It("Given a graph with S and T disconnected, when add an edge from S to T, then get nil error", func() {
			graph.AddVertex("S", "I am vertex S")
			graph.AddVertex("T", "I am vertex T")
			err := graph.AddEdge("S", "T", 10, "I am an edge")
			Expect(err).ShouldNot(HaveOccurred())
			edge, err := graph.GetEdge("S", "T")
			Expect(err).ShouldNot(HaveOccurred())
			Expect(edge.(string)).Should(BeEquivalentTo("I am an edge"))
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
			err := graph.UpdateEdgeWeight("S", "T", 10)
			Expect(err).Should(HaveOccurred())
		})

		It("Given a graph without T, when update an edge to T, then get an error", func() {
			graph.AddVertex("S", "I am vertex S")
			err := graph.UpdateEdgeWeight("S", "T", 10)
			Expect(err).Should(HaveOccurred())
		})

		It("Given a graph with S and T disconnected, when update an edge from S to T, then get nil error", func() {
			graph.AddVertex("S", "I am vertex S")
			graph.AddVertex("T", "I am vertex T")
			err := graph.UpdateEdgeWeight("S", "T", 10)
			Expect(err).Should(HaveOccurred())
		})

		It("Given a graph with S and T connected, when add an edge from S to T with -inf weight, then get an error", func() {
			graph.AddVertex("S", "I am vertex S")
			graph.AddVertex("T", "I am vertex T")
			graph.AddEdge("S", "T", 10, nil)
			err := graph.UpdateEdgeWeight("S", "T", math.Inf(-1))
			Expect(err).Should(HaveOccurred())
		})

		It("Given a graph with S and T connected, when add an edge from S to T again, then get an error", func() {
			graph.AddVertex("S", "I am vertex S")
			graph.AddVertex("T", "I am vertex T")
			graph.AddEdge("S", "T", 10, nil)
			err := graph.UpdateEdgeWeight("S", "T", 20)
			Expect(err).ShouldNot(HaveOccurred())
			edge, err := graph.GetEdgeWeight("S", "T")
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
			graph.AddVertexWithEdges(&myVertex{"S", map[ID]float64{"A": 10, "B": 10}, map[ID]float64{}})
			err := graph.CheckIntegrity()
			Expect(err).Should(HaveOccurred())
			graph.AddVertexWithEdges(&myVertex{"A", map[ID]float64{}, map[ID]float64{"S": 10, "B": 5}})
			err = graph.CheckIntegrity()
			Expect(err).Should(HaveOccurred())
			graph.AddVertexWithEdges(&myVertex{"B", map[ID]float64{"A": 5}, map[ID]float64{"S": 10}})
			err = graph.CheckIntegrity()
			Expect(err).ShouldNot(HaveOccurred())
			graph.egress["C"] = map[ID]*edge{"B": {nil, 15, true, false}}
			err = graph.CheckIntegrity()
			Expect(err).Should(HaveOccurred())
			delete(graph.egress, "C")
			graph.ingress["S"]["T"] = &edge{nil, 20, true, false}
			err = graph.CheckIntegrity()
			Expect(err).Should(HaveOccurred())
			delete(graph.ingress["S"], "T")
			graph.ingress["T"] = map[ID]*edge{"S": {nil, 20, true, false}}
			err = graph.CheckIntegrity()
			Expect(err).Should(HaveOccurred())
		})
	})

	Context("delete methods tests", func() {
		BeforeEach(func() {
			graph = NewGraph()
			graph.AddVertexWithEdges(&myVertex{"S", map[ID]float64{"A": 10, "B": 10}, map[ID]float64{}})
			graph.AddVertexWithEdges(&myVertex{"A", map[ID]float64{}, map[ID]float64{"S": 10, "B": 5}})
			graph.AddVertexWithEdges(&myVertex{"B", map[ID]float64{"A": 5}, map[ID]float64{"S": 10}})
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
			Expect(v).Should(BeEquivalentTo(&myVertex{"S", map[ID]float64{"A": 10, "B": 10}, map[ID]float64{}}))
			_, err := graph.GetVertex("S")
			Expect(err).Should(HaveOccurred())
			err = graph.CheckIntegrity()
			Expect(err).ShouldNot(HaveOccurred())
			v = graph.DeleteVertex("A")
			Expect(v).Should(BeEquivalentTo(&myVertex{"A", map[ID]float64{}, map[ID]float64{"S": 10, "B": 5}}))
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

		It("Given a graph without edge from vertex A to S, when delete edge from A to S, then nothing happens", func() {
			weight, _ := graph.GetEdgeWeight("A", "S")
			Expect(weight).Should(BeEquivalentTo(math.Inf(1)))
			edge := graph.DeleteEdge("A", "S")
			Expect(edge).Should(BeNil())
			err := graph.CheckIntegrity()
			Expect(err).ShouldNot(HaveOccurred())
		})

		It("Given a graph with edge from vertex A to  B, when delete edge from A to B, then the weight between A and B will be +inf", func() {
			weight, _ := graph.GetEdgeWeight("B", "A")
			Expect(weight).Should(BeEquivalentTo(5))
			graph.DeleteEdge("B", "A")
			weight, _ = graph.GetEdgeWeight("B", "A")
			Expect(weight).Should(BeEquivalentTo(math.Inf(1)))
			err := graph.CheckIntegrity()
			Expect(err).ShouldNot(HaveOccurred())
		})
	})

	Context("get total weight of path tests", func() {
		BeforeEach(func() {
			graph = NewGraph()
			graph.AddVertexWithEdges(&myVertex{"S", map[ID]float64{"A": 10, "B": 10}, map[ID]float64{}})
			graph.AddVertexWithEdges(&myVertex{"A", map[ID]float64{}, map[ID]float64{"S": 10, "B": 5}})
			graph.AddVertexWithEdges(&myVertex{"B", map[ID]float64{"A": 5}, map[ID]float64{"S": 10}})
		})

		AfterEach(func() {
			graph = nil
		})

		It("Given an empty path, when get its weight, then get -inf", func() {
			path := []ID{}
			Expect(graph.GetPathWeight(path)).Should(BeEquivalentTo(math.Inf(-1)))
		})

		It("Given a path contains unexisted vertex, when get its weight, then get -inf", func() {
			path := []ID{"T"}
			Expect(graph.GetPathWeight(path)).Should(BeEquivalentTo(math.Inf(-1)))
			path = []ID{"S", "A", "T"}
			Expect(graph.GetPathWeight(path)).Should(BeEquivalentTo(math.Inf(-1)))
		})

		It("Given a path only have one vertex, when get its weight, then get 0", func() {
			path := []ID{"S"}
			Expect(graph.GetPathWeight(path)).Should(BeEquivalentTo(0))
		})

		It("Given a path with vertex disconnected, when get its weight, then get +inf", func() {
			path := []ID{"B", "A", "S"}
			Expect(graph.GetPathWeight(path)).Should(BeEquivalentTo(math.Inf(1)))
		})

		It("Given a path with all vertex connected, when get its weight, then get the end to end weight of the path", func() {
			path := []ID{"S", "B", "A"}
			Expect(graph.GetPathWeight(path)).Should(BeEquivalentTo(15))
		})
	})
})

type testVertex struct {
	id     ID
	vertex Vertex
}

func (test *testVertex) ID() ID {
	return test.id
}

func (test *testVertex) Edges() (edges []Edge) {
	return test.vertex.Edges()
}
