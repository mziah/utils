// Author Masum Z. Hasan

package main

import (
	"fmt"
)

/*
	  0   1   2   3   4   5   6   7   8   9   10   11   12   13   14

  0   0   4   7   0   0   0   0   0   0   0   0    0    0    0    0

  1   4   0   0   2   4   0   0   0   0   0   0    0    0    0    0

  2   7   0   0   0   8   0   0   0   0   0   0    0    0    0    0

  3   0   2   0   0   0   0   0   7   0   0    0    0    0    0   0

  4   0   4   8   0   0   9   3   0   0   0    0    0    0   0    0

  5   0   0   0   0   9   0   0   0   6   0   0    0    0    0    0

  6   0   0   0   0   3   0   0   0   0   0   0    10   0    0    0

  7   0   0   0   7   0   0   0   0   2   5   0    0    0    0    0

  8   0   0   0   0   0   0   0   2   0   0   3    2    0    0    0

  9   0   0   0   0   0   0   0   5   0   0   0    0    8    0    0

  10  0   0   0   0   0   0   0   0   3   0   0    0    5    3    0

  11  0   0   0   0   0   0   10  0   2   0   0    0    0    5    0

  12  0   0   0   0   0   0   0   0   0   8   4    0    0    0    3

  13  0   0   0   0   0   0   0   0   0   0   3    5    0    0    11

  14  0   0   0   0   0   0   0   0   0   0   0    0    3    11   0

*/

// AdjMatrix : :

/*
// Undirected
var AdjMatrix = [][]int{
	{0, 4, 7, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},

	{4, 0, 0, 2, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},

	{7, 0, 0, 0, 8, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},

	{0, 2, 0, 0, 0, 0, 0, 7, 0, 0, 0, 0, 0, 0, 0},

	{0, 4, 8, 0, 0, 9, 3, 0, 0, 0, 0, 0, 0, 0, 0},

	{0, 0, 0, 0, 9, 0, 0, 0, 6, 0, 0, 0, 0, 0, 0},

	{0, 0, 0, 0, 3, 0, 0, 0, 0, 0, 0, 10, 0, 0, 0},

	{0, 0, 0, 7, 0, 0, 0, 0, 2, 5, 0, 0, 0, 0, 0},

	{0, 0, 0, 0, 0, 0, 6, 2, 0, 0, 3, 2, 0, 0, 0},

	{0, 0, 0, 0, 0, 0, 0, 5, 0, 0, 0, 0, 0, 0, 0},

	{0, 0, 0, 0, 0, 0, 0, 0, 3, 0, 0, 0, 0, 0, 0},

	{0, 0, 0, 0, 0, 0, 10, 0, 2, 0, 0, 0, 0, 0, 0},

	{0, 0, 0, 0, 0, 0, 0, 0, 0, 8, 4, 0, 0, 0, 3},

	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 3, 5, 0, 0, 0},

	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 11, 0}}

*/

/*
// Directed
var AdjMatrix = [][]int{
	{0, 4, 7, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},

	{0, 0, 0, 2, 4, 0, 0, 0, 0, 2, 2, 0, 0, 0, 0, 0, 0},

	{0, 0, 0, 0, 8, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},

	{0, 0, 0, 0, 0, 0, 0, 7, 0, 0, 0, 0, 0, 0, 0, 0, 0},

	{0, 0, 0, 0, 0, 9, 3, 0, 5, 3, 0, 0, 0, 0, 0, 7, 0},

	{0, 0, 0, 0, 0, 0, 0, 0, 6, 0, 0, 0, 0, 0, 0, 0, 0},

	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 10, 0, 0, 0, 0, 0},

	{0, 0, 0, 0, 0, 0, 0, 0, 4, 5, 0, 0, 0, 0, 0, 0, 0},

	{0, 3, 0, 0, 0, 0, 0, 0, 0, 0, 5, 2, 0, 0, 0, 6, 4},

	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 8, 0, 0, 0, 0},

	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 5, 0, 0, 8, 0},

	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},

	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},

	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 3, 5, 0, 0, 0, 0, 0},

	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 3, 11, 0, 0, 0},
	{0, 5, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 3, 11, 0, 5, 0},
	{0, 0, 6, 0, 0, 9, 0, 0, 0, 0, 3, 0, 0, 0, 0, 5, 0}}
*/

// row 8 and 0 swapped for node 8 to be source
var AdjMatrix1 = [][]int{
	{0, 3, 0, 0, 0, 0, 0, 0, 0, 0, 5, 2, 0, 0, 0, 6, 4},

	{3, 0, 0, 2, 4, 0, 0, 0, 0, 2, 2, 0, 0, 0, 0, 0, 0},

	{0, 0, 0, 0, 8, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},

	{0, 0, 0, 0, 0, 0, 0, 7, 0, 0, 0, 0, 0, 0, 0, 0, 0},

	{0, 0, 0, 0, 0, 9, 3, 0, 5, 3, 0, 0, 0, 0, 0, 7, 0},

	{0, 0, 0, 0, 0, 0, 7, 0, 6, 0, 0, 0, 0, 0, 0, 0, 0},

	{0, 0, 0, 0, 0, 7, 0, 0, 0, 0, 0, 10, 0, 0, 0, 0, 0},

	{0, 0, 0, 0, 0, 0, 0, 0, 4, 5, 0, 0, 0, 0, 0, 0, 5},

	{0, 4, 7, 0, 12, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},

	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 8, 0, 0, 0, 0},

	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 5, 0, 0, 8, 0},

	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},

	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},

	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 3, 5, 0, 0, 0, 0, 0},

	{0, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, 0, 3, 11, 0, 0, 0},
	{0, 5, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 3, 11, 0, 5, 0},
	{0, 0, 6, 0, 0, 9, 0, 0, 0, 0, 3, 0, 0, 0, 0, 5, 0}}

// Complete graph
var AdjMatrix2 = [][]int{
	{4, 3, 8, 2, 7, 8},
	{4, 3, 8, 2, 7, 8},
	{4, 3, 5, 2, 7, 8},
	{4, 3, 8, 2, 9, 8},
	{4, 2, 8, 2, 7, 8},
	{4, 3, 1, 2, 4, 8}}

// Tree
var AdjMatrix3 = [][]int{
	{0, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},

	{0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},

	{0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0},

	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},

	{0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0},

	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},

	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},

	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},

	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1},

	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0},

	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},

	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},

	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},

	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},

	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}}

// NUMNODES : :
var NUMNODES int

// NUMEDGES : :
var NUMEDGES int

// Edge : :
type Edge struct {
	id     int
	src    int  // Node id
	dst    int  // Node id
	params int  // can be multivariate like BW, delay, jitter and/or colors; we'll use it as weight
	spt    bool // Is this edge part of SPT?
}

// Node : :
type Node struct {
	id      int
	aedges  []int // ids array of adjacent edges
	nedges  int   // # of adjacent edges
	params  int   // Node weight
	visited bool
}

// Graph : :
type Graph struct {
	nodes []Node
	edges []Edge
}

// Stacki : : Stack with integer
type Stacki struct {
	value []int // Will contain index to Nodes
	count int
}

func (s *Stacki) Push(value int) {

	// Memory allocated fixed amount (in main) when Stacki created
	// Increase slice ([]value) size 3 times (in anticipation of stack growth), when count increases and more value pushed
	if s.count > len(s.value) {
		v := make([]int, len(s.value)*3)
		// copy existing stack value to newly allocated memory pointed to by v
		copy(v, s.value)
		// It now point to new memory
		s.value = v
	}
	// Push value into stack
	s.value[s.count] = value
	s.count++

}

func (s *Stacki) Pop() int {
	if s.count == 0 {
		return 0
	}
	v := s.value[s.count-1]
	s.count--
	return v
}

// FifoQ : :
type FifoQi struct {
	value []int // Will contain index to Nodes
	head  int
	tail  int
	count int
}

func (q *FifoQi) Push(value int) {

	// Memory allocated fixed amount (in main) when FifoQi created
	// Increase slice ([]value) size 3 times (in anticipation of Q growth),
	// when count increases beyond current Q length and more value pushed
	if q.count == len(q.value) {
		v := make([]int, len(q.value)*3)
		// copy existing Q contents from head to end to newly allocated memory pointed to by v
		copy(v, q.value[q.head:])
		// It now points to new memory
		q.value = v
		q.head = 0
		q.tail = len(q.value)
	}
	// Add value at tail
	q.value[q.tail] = value
	//q.tail = (q.tail + 1) % len(q.nodes)
	q.tail = q.tail + 1

	q.count++
}

func (q *FifoQi) Pop() int {
	if q.count == 0 {
		return -1
	}
	v := q.value[q.head]
	//q.head = (q.head + 1) % len(q.value)
	q.head = q.head + 1
	q.count--
	return v
}

// Total # of edges in the graph
func totalEdges(input [][]int) int {
	var numEdges = 0
	for i := 0; i < NUMNODES; i++ {
		for j := 0; j < NUMNODES; j++ {
			if input[i][j] != 0 {
				numEdges++
			}
		}
	}
	return numEdges
}

func generateDOTGR(gr *Graph) {

	fmt.Printf("digraph G { \n")
	for i := range gr.nodes {
		for k := range gr.nodes[i].aedges {
			edgeid := gr.nodes[i].aedges[k]
			p := gr.edges[edgeid].params
			dst := gr.edges[edgeid].dst
			fmt.Printf("%d -> %d[label=\"%d\", weight=\"%d\"];\n", i, dst, p, p)
		}
	}
	fmt.Printf("}\n")
}

// # of edges in a given node
func numEdgesInNode(in [][]int, i int) int {
	numEdges := 0
	for j := 0; j < NUMNODES; j++ {
		if in[i][j] != 0 {
			numEdges++
		}
	}
	return numEdges
}

// Given an adjacency matrix representation of graph, initialize struct representation of the graph
//func initGraph(input [][]int) *Graph {
func initGraph(input [][]int) *Graph {

	gr := new(Graph)
	gr.nodes = make([]Node, NUMNODES)
	gr.edges = make([]Edge, NUMEDGES)
	// Create Nodes
	for i := 0; i < NUMNODES; i++ {
		gr.nodes[i].id = i
		gr.nodes[i].params = 1
		ne := numEdgesInNode(input, i)
		gr.nodes[i].aedges = make([]int, ne)
		gr.nodes[i].nedges = ne
		gr.nodes[i].visited = false
	}
	// Create Edges; All Nodes should be created already
	// l: for all the edges in graph
	l := 0
	// n: edges per node
	n := 0
	for k := 0; k < NUMNODES; k++ {

		for j := 0; j < NUMNODES; j++ {
			if input[k][j] != 0 {

				// Edges attached to nodes
				gr.edges[l].id = l
				gr.edges[l].src = k
				gr.edges[l].dst = j
				gr.edges[l].params = input[k][j]
				gr.nodes[k].aedges[n] = l

				n++
				l++
			}
		}
		n = 0
	}
	return gr
}

//Breadth First Traversal and print graph
func dfTraversal(gr *Graph, root int) {

	if !gr.nodes[root].visited {
		gr.nodes[root].visited = true
		for i := range gr.nodes[root].aedges {
			eid := gr.nodes[root].aedges[i]
			p := gr.edges[eid].params
			dst := gr.edges[eid].dst
			// Edges are placed onto a FIFO Q so that nodes are visited breath first
			// For Depth first each function call is pushed into stack by the (language) execution environment,
			// hence explicit stack manipulaton is not needed

			fmt.Printf("%d -> %d[label=\"%d\", weight=\"%d\"];\n", root, dst, p, p)
			dfTraversal(gr, dst)
		}
	}
}

func main() {

	NUMNODES = len(AdjMatrix1)
	NUMEDGES = totalEdges(AdjMatrix1)
	gr := initGraph(AdjMatrix1)
	source := 0

	fmt.Printf("\n\n")
	fmt.Printf("***** Graph 1 ***** \n\n")
	generateDOTGR(gr)
	fmt.Printf("\n\n")

	fmt.Printf("***** DF Traversal 1 ***** \n\n")
	dfTraversal(gr, source)

	NUMNODES = len(AdjMatrix2)
	NUMEDGES = totalEdges(AdjMatrix2)
	gr = initGraph(AdjMatrix2)
	source = 0

	fmt.Printf("\n\n")
	fmt.Printf("***** Graph 2 ***** \n\n")
	generateDOTGR(gr)
	fmt.Printf("\n\n")

	fmt.Printf("***** DF Traversal 2 ***** \n\n")
	dfTraversal(gr, source)

	NUMNODES = len(AdjMatrix3)
	NUMEDGES = totalEdges(AdjMatrix3)
	gr = initGraph(AdjMatrix3)
	source = 0

	fmt.Printf("\n\n")
	fmt.Printf("***** Graph 3 ***** \n\n")
	generateDOTGR(gr)
	fmt.Printf("\n\n")

	fmt.Printf("***** DF Traversal 3 ***** \n\n")
	dfTraversal(gr, source)
}
