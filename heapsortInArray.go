/* Heap Sort
 * In place sort - sorted data is in same input array (list)
 * Masum Z. Hasan
 */

package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

// ListSize : :
var ListSize int

// Binary Min Heap is constructed on the same input array
// CurrentSize : : Used to keep track of current size of Heap as its size
// keeps decrementing everytime an element is moved at the end oth list
var CurrentSize int

const HEAPFILE = "./dot/heap"

// Mean heap file count
var MHC = 0

var SOURCE int

func createMinHeap(list []int, rootIdx int) {

	printMinHeap(list)

	var temp int
	idxOfSmallerChild := rootIdx
	// In a binary tree at level l there are 2^l (^  : power) elements
	// If index of 1st element at level l is n, then last element index will be 2*n
	// Hence index of children at level l+1 of an index (root) n at level l will be 2*n+1 (left child)
	// and 2*n+2 (right child)
	leftIdx := 2*rootIdx + 1
	rightIdx := 2*rootIdx + 2

	// leftIdx or rightIdx < listSize check: any index (node) > list size will not have child
	// (as element at that index non-existent)
	// Will check nil if represented as struct/graph
	if leftIdx < CurrentSize && list[leftIdx] < list[idxOfSmallerChild] {
		idxOfSmallerChild = leftIdx
	}
	if rightIdx < CurrentSize && list[rightIdx] < list[idxOfSmallerChild] {
		idxOfSmallerChild = rightIdx
	}

	// Index (idxOfSmallerChild) of smaller child identified
	// Now swap this child with parent, if they are not the same

	if idxOfSmallerChild != rootIdx {

		temp = list[rootIdx]
		list[rootIdx] = list[idxOfSmallerChild]
		list[idxOfSmallerChild] = temp

		// Debug
		//printMinHeap(list, rootIdx, idxOfSmallerChild)

		// Now resursively push parent down the heap and smaller element up the heap
		createMinHeap(list, idxOfSmallerChild)
	}
}

// Extract min, adjust heap -- list is sorted if called listSize-1 times
func extractMinAdjust(list []int) int {

	// Creste Min Heap
	// listSize/2 -1 are all the non-leaf nodes
	// Start with lowest level of non-leaf nodes and their children
	// Move smaller values up, push larger values down

	for i := CurrentSize/2 - 1; i >= 0; i-- {
		createMinHeap(list, i)
	}

	// All roots now less than its children, but not yet sorted

	// Debug
	//printMinHeap(list, len(list))

	min := list[0]

	// swap last element (list[CurrentSize-1] and root (list[0])
	list[0] = list[CurrentSize-1]
	list[CurrentSize-1] = min

	//fmt.Printf("last and root swapped\n")
	//printMinHeap(list, len(list))

	CurrentSize--

	return min
}

func heapSort(list []int) {

	// Create min heap

	for i := len(list)/2 - 1; i >= 0; i-- {
		createMinHeap(list, i)
	}
	// All roots now less than its children, but not yet sorted

	/* Debug */

	/* Sort the min-heap: exchange top root with the last element of the heap reducing heap (list) size by 1 every time */
	/* Then apply createMinHeap on the top root, thus recursively pushing down value in top root and moving up next smaller value to root */

	for i := len(list) - 1; i >= 0; i-- {
		/* swap root and elemnt at i */
		temp := list[0]
		list[0] = list[i]
		list[i] = temp

		/* Call createMinHeap on root to push down root or bring next min to root */
		createMinHeap(list, 0)
	}

}

func printMinHeap(list []int) {
	//fmt.Printf("Min-heap:\n")

	fl := HEAPFILE + "-" + strconv.Itoa(SOURCE) + "-" + strconv.Itoa(MHC) + ".dot"
	line := fmt.Sprintf("graph {\n")
	print2File(fl, line)

	//listSize := len(list)
	listSize := len(list)
	for i := 0; i < listSize; i++ {
		leftIdx := 2*i + 1
		rightIdx := 2*i + 2
		if leftIdx < listSize {
			// print Min Heap in Graphviz Dot format
			line = fmt.Sprintf("%d -- %d;\n", list[i], list[leftIdx])
			print2File(fl, line)
		}
		if rightIdx < listSize {
			line = fmt.Sprintf("%d -- %d;\n", list[i], list[rightIdx])
			print2File(fl, line)
		}
	}
	//fmt.Printf("\n\n")
	line = fmt.Sprintf("}\n")
	print2File(fl, line)
	MHC++
}

func print2File(finename, line string) {

	f, err := os.OpenFile(finename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := f.Write([]byte(line)); err != nil {
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}

func main() {

	list := []int{27, 54, 18, 9, 72, 90, 81, 5, 92, 43, 88, 99}
	//list3 := []int{5, 9, 18, 27, 43, 54, 72, 81, 88, 90, 92, 99}
	//list := []int{99, 92, 90, 88, 81, 72, 54, 43, 27, 18, 9, 5}
	//list := []int{0, 9999, 9999, 9999, 9999, 9999, 9999, 9999, 9999, 9999, 9999, 9999}

	ListSize := len(list)
	fmt.Printf("Unsorted List = ")
	for i := 0; i < ListSize; i++ {
		fmt.Printf("%d ", list[i])
	}
	fmt.Printf("\n\n")

	// SHOULD NOT PASS GLOBAL VARIABLE IN FUNC - becomes pass by value
	CurrentSize = ListSize

	// Extract Min and adjust heap - list is sorted when loop ends or call heapSort

	//fmt.Printf("Sorted list = ")

	for i := 0; i < ListSize; i++ {

		printMinHeap(list)
		// _ : if return value not used
		_ = extractMinAdjust(list)
		//min := extractMinAdjust(list)
		//fmt.Printf("%d ", min)
		printMinHeap(list)
	}

	fmt.Printf("\n\nSorted List Decreasing order = ")
	for i := 0; i < ListSize; i++ {
		fmt.Printf("%d ", list[i])
	}
	fmt.Printf("\n\n")

	//list2 := make([]int, len(list))
	fmt.Printf("\n\nSorted List increasing order = ")
	for i := ListSize - 1; i >= 0; i-- {
		fmt.Printf("%d ", list[i])
	}

	fmt.Printf("\n\n")

}
