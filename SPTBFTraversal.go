//Author: Masum Z. Hasan

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
var AdjMatrix = [][]int{
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

/*
// Complete graph
var AdjMatrix = [][]int{
	{4, 3, 8, 2, 7, 8},
	{4, 3, 8, 2, 7, 8},
	{4, 3, 5, 2, 7, 8},
	{4, 3, 8, 2, 9, 8},
	{4, 2, 8, 2, 7, 8},
	{4, 3, 1, 2, 4, 8}}
*/

// BIGNUMBER : : For Shortest path weight calc
const BIGNUMBER = 99999999

const NECMP = 10 // # of ECMP supported

// NUMNODES : :
//const NUMNODES = 17
var NUMNODES = len(AdjMatrix)

// NUMEDGES : :
var NUMEDGES = totalEdges(AdjMatrix)

//const NUMNODES = 6

// NodesVisited : : Used in Breadth first traversal to exit recursion
var NodesVisited int

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
	id       int
	aedges   []int      // ids array of adjacent edges
	nedges   int        // # of adjacent edges
	params   int        // sum of shortest path weight from source to this node
	lpr      bool       // Previous lower param replaced - to keep track of ECMP
	sptedge  [NECMP]int // Id of the edge leading to shortest path weight; Array for 10 ECMP
	nsptedge int        // current count of sptedge
	visited  bool
	addedMH  bool // Added to MinHeapSet
	onspt    bool // This node is on shortest path
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
		return 0
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

// GraphViz DOT notation to visualize graph
func generateDOTAM() {

	fmt.Printf("digraph G { \n")
	for i := 0; i < NUMNODES; i++ {
		for j := 0; j < NUMNODES; j++ {
			if AdjMatrix[i][j] != 0 {
				fmt.Printf("%d -> %d[label=\"%d\", weight=\"%d\"];\n", i, j, AdjMatrix[i][j], AdjMatrix[i][j])
			}
		}
	}
	fmt.Printf("}\n")
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

// # if edges in a given node
func numEdgesInNode(i int) int {
	numEdges := 0
	for j := 0; j < NUMNODES; j++ {
		if AdjMatrix[i][j] != 0 {
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
		gr.nodes[i].params = BIGNUMBER
		ne := numEdgesInNode(i)
		gr.nodes[i].aedges = make([]int, ne)
		gr.nodes[i].nedges = ne
		gr.nodes[i].visited = false
		//gr.nodes[i].sptedge = BIGNUMBER
		gr.nodes[i].addedMH = false
		gr.nodes[i].lpr = false
		gr.nodes[i].onspt = false
		gr.nodes[i].nsptedge = 0
		for j := 0; j < NECMP; j++ {
			gr.nodes[i].sptedge[j] = BIGNUMBER
		}

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
				gr.edges[l].spt = false
				gr.nodes[k].aedges[n] = l

				n++
				l++
			}
		}
		n = 0
	}
	return gr
}

func calcShortestPathWeight(gr *Graph, root int, j int) int {
	// id of the edge
	eid := gr.nodes[root].aedges[j]
	// Destination id or index of the edge
	dst := gr.edges[eid].dst
	// src node weight + edge weight
	p := gr.nodes[root].params + gr.edges[eid].params
	// min (src node weight + edge weight, dst node weight)
	if p < gr.nodes[dst].params {
		gr.nodes[dst].params = p
		// Arrived to dst via eid
		gr.nodes[dst].sptedge[0] = eid
		gr.edges[eid].spt = true
	}
	return dst
}

// Recursive - Breadth first Traversal with FIFO Q (No Min Heap)
// Traverses all the edges
// For example: 0 16 5 & 0 16 2 4 5 & 0 1 3 7 8 2 4 5
// With Min Heap it should traverse only the 1st path
func calcShortestPathTreeBF(gr *Graph, root int, fifoq *FifoQi) {

	for NodesVisited < NUMNODES {
		if !gr.nodes[root].visited {
			gr.nodes[root].visited = true
			for j := range gr.nodes[root].aedges {
				dst := calcShortestPathWeight(gr, root, j)
				// Nodes are placed onto a FIFO Q so that nodes are visited breadth first
				// For Depth first each function call is pushed into stack by the execution environment,
				// hence explicit stack manipulaton is not needed

				fifoq.Push(dst)
			}
			NodesVisited++
		}
		calcShortestPathTreeBF(gr, fifoq.Pop(), fifoq)
	}
}

func printSPT(gr *Graph) {

	// Prints in Graphviz DOT notation
	for i := 0; i < NUMNODES; i++ {
		for j := range gr.nodes[i].aedges {
			eid := gr.nodes[i].aedges[j]
			dst := gr.edges[eid].dst
			fmt.Printf("%d[label=\"%d, %d\"]\n", dst, dst, gr.nodes[dst].params)
			if gr.nodes[dst].sptedge[0] == eid {
				fmt.Printf("%d -> %d[label=\"%d\", weight=\"%d\", color=red];\n", gr.nodes[i].id, dst, gr.edges[eid].params, gr.edges[eid].params)
			} else {
				fmt.Printf("%d -> %d[label=\"%d\", weight=\"%d\"];\n", gr.nodes[i].id, dst, gr.edges[eid].params, gr.edges[eid].params)
			}
		}
		fmt.Printf("\n")
	}
}

func main() {

	//numEdges := totalEdges(AdjMatrix)

	gr := initGraph(AdjMatrix)
	source := 7
	NodesVisited = 0
	gr.nodes[source].params = 0
	gr.nodes[source].sptedge[0] = 0

	fifoq1 := &FifoQi{value: make([]int, NUMEDGES)}
	fmt.Printf("***** Breadth First Traversal ***** \n\n")
	calcShortestPathTreeBF(gr, source, fifoq1)
	fifoq1 = nil

	fmt.Printf("\n Shortest Path Tree - BF Traversal - colored red from Source = %d\n\n", source)
	fmt.Println("use https://dreampuf.github.io/GraphvizOnline/ and paste DOT output to render graph")
	fmt.Println("digraph G {")
	printSPT(gr)
	fmt.Println("}")

}
