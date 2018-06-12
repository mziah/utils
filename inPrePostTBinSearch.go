package main

import (
	"fmt"
)

/* Heap Sort
 *
 * Masum Z. Hasan
 *
 *
 * List (Array) to be sorted {27, 54, 18, 9, 72, 90, 81, 5, 92, 43, 88, 99} represented as complete binary tree
 * listSize = 12
 *
 *                                         27                    Level 0 2^0 (1) max nodes
 *                                     /        \
 *                                   54          18              Level 1 2^1 (2) max nodes
 *                                 /   \       /   \
 *                                9     72    90    81           Level 2 2^2 (4) max nodes
 *                               / \   / \   / \   /  \
 *                              5  92 43 88  99                  Level 3 2^3 (8) max nodes but fewer in this example
 *
 *
 * We will create Max-heap (as opposed to Min-heap): every root is greater than its children
 * Each root  of subgraph (having child[ren]) will have the Max-heap property
 * All the non-leave nodes "live" in indexexindices listSize / 2 -1 (index starts from 0)
 * Hence we start at listSize / 2 -1 : in this case 12/2 -1 = 5, that is index 0-5 (numbers 27 to 90)
 * We go backward from index=5 to index=0
 *
 * leftIdx/rightIdx < listSize checks whether a subgraph node has left or right child
 *
 *
 *                                         27
 *                                     /        \
 *                                   54          18
 *                                 /   \       /   \        Level 2
 *                                92    88    99    81      Each Level 2 node & its larger child, if any, swapped
 *                               / \   / \   / \   /  \     Larger, for example 99, swapped with parent node containing 90
 *                              5  9  43 72  90             createMinHeap will be called recursively on child that was swapped -
 *                                                          90 in this case, but it will exit
 *
 *
 *                                         27
 *                                     /        \           Level 1
 *                                   92          99         Each Level 1 node & its larger child, if any, swapped
 *                                 /   \       /   \        Parent 18 and larger child 99 swapped
 *                                54    88    18    81      createMinHeap will be called recursively on child that was swapped -
 *                               / \   / \   / \   /  \     18 in this case
 *                              5  9  43 72  90
 *
 *
 *                                         27
 *                                     /        \           Level 1
 *                                   92          99
 *                                 /   \       /   \
 *                                54    88    90    81      18 pushed down recursively: 90 & 18 swapped
 *                               / \   / \   / \   /  \
 *                              5  9  43 72  18
 *
 *
 *
 *                                         99               Level 0
 *                                     /        \           99, 27 swapped
 *                                   92          27
 *                                 /   \       /   \
 *                                54    88    90    81
 *                               / \   / \   / \   /  \
 *                              5  9  43 72  18
 *
 *
 *
 *
 *                                         99               Level 0
 *                                     /        \
 *                                   92          90
 *                                 /   \       /   \        27 pushed down recursively
 *                                54    88    27    81
 *                               / \   / \   / \   /  \
 *                              5  9  43 72  18
 *
 *                              Min heap created; now sort in ascending order
 *                              Since ascending order - swap last element (18) with root containing max
 *                              Create max-heap again by leaving out the last element or
 *                              size = size -1
 *
 *
 *                                         92               Root (99) and last element (18) swapped
 *                                     /        \           Now call create_maxheap with root
 *                                   18          90         Next larger element 92 will be at root and
 *                                 /   \       /   \        18 pushed down recursively
 *                                54    88    27    81
 *                               / \   / \   / \   /  \
 *                              5  9  43 72  99
 *
 *
 *
 *                                         92               Root (99) and last element (18) swapped
 *                                     /        \           Now call create_maxheap with root
 *                                   88          90         Next larger element 92 will be at root and
 *                                 /   \       /   \        18 pushed down recursively
 *                                54    18    27    81
 *                               / \   / \   / \   /  \
 *                              5  9  43 72  99
 *
 *
 *
 *                                         72
 *                                     /        \
 *                                   88          90
 *                                 /   \       /   \
 *                                54    18    27    81
 *                               / \   / \   / \   /  \
 *                              5  9  43 92  99
 *
 *
 *
 *                                         90               Push 72 down
 *                                     /        \
 *                                   88          72
 *                                 /   \       /   \
 *                                54    18    27    81
 *                               / \   / \   / \   /  \
 *                              5  9  43 92  99
 *
 *
 *
 *                                         90               Push 72 down
 *                                     /        \
 *                                   88          81
 *                                 /   \       /   \
 *                                54    18    27    72
 *                               / \   / \   / \   /  \
 *                              5  9  43 92  99
 *
 *
 *                                         42               swap 90, 42
 *                                     /        \
 *                                   88          81
 *                                 /   \       /   \
 *                                54    18    27    72
 *                               / \   / \   / \   /  \
 *                              5  9  90 92  99
 *
 *
 *                                         88               create_maxheap -- push down 42
 *                                     /        \
 *                                   42          81
 *                                 /   \       /   \
 *                                54    18    27    72
 *                               / \   / \   / \   /  \
 *                              5  9  90 92  99
 *
 *
 *
 *                                         88               create_maxheap -- push down 42
 *                                     /        \
 *                                   54          81
 *                                 /   \       /   \
 *                                42    18    27    72
 *                               / \   / \   / \   /  \
 *                              5  9  90 92  99
 *
 *
 *
 *                                         9               swap 88, 9
 *                                     /        \
 *                                   54          81
 *                                 /   \       /   \
 *                                42    18    27    72
 *                               / \   / \   / \   /  \
 *                              5  88 90 92  99
 *
 *
 *                                         81               create_maxheap -- push down 9
 *                                     /        \
 *                                   54          9
 *                                 /   \       /   \
 *                                42    18    27    72
 *                               / \   / \   / \   /  \
 *                              5  88 90 92  99
 *
 *
 *                                         81               create_maxheap -- push down 9
 *                                     /        \
 *                                   54          72
 *                                 /   \       /   \
 *                                42    18    27    9
 *                               / \   / \   / \   /  \
 *                              5  88 90 92  99
 *
 *
 *                                         5               swap 81, 5
 *                                     /        \
 *                                   54          72
 *                                 /   \       /   \
 *                                42    18    27    9
 *                               / \   / \    / \  /  \
 *                              81  88 90 92  99
 *
 *
 *
 *                                         72               create_maxheap -- push down 5
 *                                     /        \
 *                                   54          5
 *                                 /   \       /   \
 *                                42    18    27    9
 *                               / \   / \    / \  /  \
 *                              81  88 90 92  99
 *
 *
 *
 *                                         72               create_maxheap -- push down 5
 *                                     /        \
 *                                   54          27
 *                                 /   \       /   \
 *                                42    18     5    9
 *                               / \   / \    / \  /  \
 *                              81  88 90 92  99
 *
 *
 *
 *                                         9               swap 72, 9
 *                                     /        \
 *                                   54          27
 *                                 /   \       /   \
 *                                42    18     5    72
 *                               / \   / \    / \  /  \
 *                              81  88 90 92  99
 *
 *
 *
 *                                         54               create_maxheap -- push down 9
 *                                     /        \
 *                                   9          27
 *                                 /   \       /   \
 *                                42    18     5    72
 *                               / \   / \    / \  /  \
 *                              81  88 90 92  99
 *
 *
 *                                         54               create_maxheap -- push down 9
 *                                     /        \
 *                                   42         27
 *                                 /   \       /   \
 *                                9    18     5    72
 *                               / \   / \    / \  /  \
 *                              81  88 90 92  99
 *
 *
 *                                         42               swap 54, 5 create_maxheap -- push down 5
 *                                     /        \
 *                                   18         27
 *                                 /   \       /   \
 *                                9    5      54    72
 *                               / \   / \    / \  /  \
 *                              81  88 90 92  99
 *
 *
 *
 *                                         27               swap 42, 5 create_maxheap -- push down 5
 *                                     /       \
 *                                   18         5
 *                                 /   \       /   \
 *                                9    42     54    72
 *                               / \   / \    / \  /  \
 *                              81  88 90 92  99
 *
 *
 *                                         18               swap 27, 9 create_maxheap -- push down 9
 *                                     /        \
 *                                   9           5
 *                                 /   \       /   \
 *                                27    42    54    72
 *                               / \   / \    / \  /  \
 *                              81  88 90 92  99
 *
 *
 *
 *                                         5               swap 18, 5 create_maxheap -- push down 5
 *                                     /        \
 *                                   9           18
 *                                 /   \       /   \
 *                                27    42    54    72
 *                               / \   / \    / \  /  \
 *                              81  88 90 92  99
 *
 *
 *                                         9               swap 18, 5 create_maxheap -- push down 5
 *                                     /        \
 *                                   5           18
 *                                 /   \       /   \
 *                                27    42    54    72
 *                               / \   / \    / \  /  \
 *                              81  88 90 92  99
 *
 *
 *
 *                                         5               swap 9, 5 create_maxheap entered with size=0; quits
 *                                     /        \
 *                                   9           18
 *                                 /   \       /   \
 *                                27    42    54    72
 *                               / \   / \    / \  /  \
 *                              81  88 90 92  99
 *
 * Compoutational and space complexity
 *
 * Space complexity O(1) as sorting is done in situ
 *
 * An element can be pushed down from root to leaf
 * The height of a binary tree of size is log base 2 n
 * for Max heap creation that can be n/2*logn in the worst case
 * For sorting again that can be n*logn
 * Hence computational complexity is O(nlogn) in average and worst case
 * For list [] = {27,54,18,9,72,90,81,5,92,43,88,99} create_maxheap was eneterd 43 times
 * For list [] = 5 9 18 27 43 54 72 81 88 90 92 99, create_maxheap was eneterd 47 times
 * For list [] = {99,92,90,88,81,72,54,43,27,18,9,5}; create_maxheap was eneterd 35 times (this list is already max-heap)
 ******************/

// Recursively called to create min heap: root of a subgraph is smaller than its children
// Smaller child (c) is swapped with its parent (p)
// Then p, which is now at the position of c is pushed down recursively

// CurrentSize : :
var CurrentSize = 0

type BST struct {
	value int
	left  *BST
	right *BST
}

func (root *BST) insert(v int) {

	var lr bool
	var rp *BST
	r := root

	fmt.Printf("Trace Traversal: Insert %d into BST\n", v)

	for r != nil {
		if v <= r.value {
			// rp: parent to insert v
			rp = r
			fmt.Printf("left  %d  ", rp.value)
			r = r.left
			//r.insert(v)
			// left or right parent to insert to
			lr = true

		} else {
			rp = r
			fmt.Printf("right  %d  ", rp.value)

			r = r.right

			//r.insert(v)
			lr = false

		}
	}
	var child *BST
	child = new(BST)
	child.value = v
	child.left = nil
	child.right = nil
	if lr {
		rp.left = child
	} else {
		rp.right = child
	}

}

// Create BST from existing sorted array (list)
func createBSTFromList(list []int, left, right int) *BST {

	var middle int
	var root *BST

	// leaf nodes
	// left: left to middle - 1 (right): middle - 1 will keep moving left going past left,
	// thus becoming > right
	// right: middle + 1, which is left for right subgraph, which will keep moving right
	// going past right, thus becoming > right
	if left > right {
		return nil // root.left/right will be nil for leaf nodes
	} else {
		middle = (left + right) / 2
		root = new(BST)
		root.value = list[middle]
		root.left = createBSTFromList(list, left, middle-1)
		root.right = createBSTFromList(list, middle+1, right)

		return root
	}
}

// Inorder traversal and print
// Left subgraph --> root --> right subgraph
func inOrderBST(root *BST) {

	if root != nil {
		inOrderBST(root.left)
		fmt.Printf(" %d ", root.value)
		inOrderBST(root.right)
	}
}

// Preorder traversal and print
// Root --> Left subgraph --> right subgraph
func preOrderBST(root *BST) {

	if root != nil {
		fmt.Printf(" %d ", root.value)
		preOrderBST(root.left)
		preOrderBST(root.right)
	}
}

// Post traversal and print
// Left subgraph --> right subgraph --> root
func postOrderBST(root *BST) {

	if root != nil {
		postOrderBST(root.left)
		postOrderBST(root.right)
		fmt.Printf(" %d ", root.value)
	}
}

func searchBST(node *BST, find int) bool {
	found := false
	n := node
	for n != nil {
		if find == n.value {
			found = true
			break
		} else {
			if find < n.value {
				//fmt.Printf("v  %d  ", n.value)
				n = n.left
			} else {
				//fmt.Printf("v  %d  ", n.value)
				n = n.right
			}
		}
	}
	return found
}

func (r *BST) searchBST2(find int) bool {
	found := false
	for r != nil {
		if find == r.value {
			found = true
			break
		} else {
			if find < r.value {
				fmt.Printf("\n Value left  %d", r.value)
				r = r.left
			} else {
				fmt.Printf("\n Value right  %d", r.value)

				r = r.right
			}
		}
	}
	fmt.Println()
	return found
}

func printBSTDOT(root *BST) {

	if root != nil {
		fmt.Printf("%d [label=%d];\n", root.value, root.value)
		if root.left != nil {
			fmt.Printf("%d [label=%d];\n", root.left.value, root.left.value)
			fmt.Printf("%d -- %d;\n", root.value, root.left.value)
		}
		if root.right != nil {
			fmt.Printf("%d [label=%d];\n", root.right.value, root.right.value)
			fmt.Printf("%d -- %d;\n", root.value, root.right.value)
		}
		printBSTDOT(root.left)
		printBSTDOT(root.right)
	}
}
func createMinHeap(list []int, rootIdx, listSize int) {

	var temp int
	idxOfSmallerChild := rootIdx
	// In a binary tree at level l there are 2^l (^  : power) elements
	// If index of 1st element at level l is n, then last element index will be 2*n
	// Hence index of children at level l+1 of an index (root) n at level l will be 2*n+1 (left child) and 2*n+2 (right child)
	leftIdx := 2*rootIdx + 1
	rightIdx := 2*rootIdx + 2

	// leftIdx or rightIdx < listSize check: any index (node) > list size will not have child (as element at that index non-existent)
	// Will check nil if represented as struct/graph
	if leftIdx < listSize && list[leftIdx] < list[idxOfSmallerChild] {
		idxOfSmallerChild = leftIdx
	}
	if rightIdx < listSize && list[rightIdx] < list[idxOfSmallerChild] {
		idxOfSmallerChild = rightIdx
	}

	// Index (idxOfSmallerChild) of smaller child identified
	// Now swap this child with parent, if they are not the same

	if idxOfSmallerChild != rootIdx {

		temp = list[rootIdx]
		list[rootIdx] = list[idxOfSmallerChild]
		list[idxOfSmallerChild] = temp

		// Debug
		// printMinHeap(list, listSize)

		// Now resursively push parent down the heap
		createMinHeap(list, idxOfSmallerChild, listSize)
	}
}

func printMinHeap(list []int, listSize int) {
	fmt.Printf("Min-heap:\n")

	for i := 0; i < listSize; i++ {

		leftIdx := 2*i + 1
		rightIdx := 2*i + 2
		if leftIdx < listSize {
			// print Min Heap in Graphviz Dot format
			fmt.Printf("%d -- %d;\n", list[i], list[leftIdx])
		}
		if rightIdx < listSize {
			fmt.Printf("%d -- %d;\n", list[i], list[rightIdx])
		}
	}
	fmt.Printf("\n\n")
}

// Extract min, adjust heap -- list is sorted if called listSize-1 times
func extractMinAdjust(list []int) int {

	// Creste Min Heap
	// listSize/2 -1 are all the non-leaf nodes
	// Start with lowest level of non-leaf nodes and their children
	// Move smaller values up, push larger values down

	for i := CurrentSize/2 - 1; i >= 0; i-- {
		createMinHeap(list, i, CurrentSize)
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

func heapSort(list []int, listSize int) {

	// Create min heap

	for i := listSize/2 - 1; i >= 0; i-- {
		createMinHeap(list, i, listSize)
	}
	// All roots now less than its children, but not yet sorted

	/* Debug */

	/* Sort the min-heap: exchange top root with the last element of the heap reducing heap (list) size by 1 every time */
	/* Then apply createMinHeap on the top root, thus recursively pushing down value in top root and moving up next smaller value to root */

	for i := listSize - 1; i >= 0; i-- {
		/* swap root and elemnt at i */
		temp := list[0]
		list[0] = list[i]
		list[i] = temp

		/* Call createMinHeap on root to push down root or bring next min to root */
		createMinHeap(list, 0, i)
	}

}

func main() {

	list := []int{27, 54, 18, 9, 72, 90, 81, 5, 92, 43, 88, 99}
	list2 := []int{27, 54, 18, 9, 72, 90, 81, 5, 92, 43, 88, 99}

	list3 := []int{5, 9, 18, 27, 43, 54, 72, 81, 88, 90, 92, 99}
	//list := []int{99, 92, 90, 88, 81, 72, 54, 43, 27, 18, 9, 5}
	//list := []int{0, 9999, 9999, 9999, 9999, 9999, 9999, 9999, 9999, 9999, 9999, 9999}

	size := len(list)
	fmt.Printf("\nUnsorted List = ")
	for i := 0; i < size; i++ {
		fmt.Printf("%d ", list[i])
	}
	fmt.Println()

	// printMinHeap(list, size)

	// SHOULD NOT PASS GLOBAL VARIABLE IN FUNC - becomes pass by value
	CurrentSize = size

	// Extract Min and adjust heap - list is sorted when loop ends or call heapSort
	for i := 0; i < size; i++ {

		_ = extractMinAdjust(list)
		//fmt.Printf("Min = %d\n", min1)
	}

	fmt.Printf("\nSorted List Decreasing order = ")
	for i := 0; i < size; i++ {
		fmt.Printf("%d ", list[i])
	}
	fmt.Println()

	//list2 := make([]int, len(list))
	fmt.Printf("\nSorted List increasing order = ")
	for i := size - 1; i >= 0; i-- {
		fmt.Printf("%d ", list[i])
	}
	fmt.Println()

	fmt.Printf("\nBST\n")

	// Create BST from list directly
	root := createBSTFromList(list3, 0, len(list3)-1)

	fmt.Println("diagraph G {")
	printBSTDOT(root)
	fmt.Println("}")

	fmt.Printf("\nIn/pre/post order traversal:\n")

	inOrderBST(root)
	fmt.Printf("\n\n")
	preOrderBST(root)
	fmt.Printf("\n\n")
	postOrderBST(root)

	fmt.Printf("\n\nFind 9 in BST:")

	fmt.Printf("Result: %t\n\n", root.searchBST2(9))

	// Create BST by inserting from list
	root2 := new(BST)
	root2.value = list2[0]
	root2.left = nil
	root2.right = nil
	for i := 1; i < size; i++ {
		root2.insert(list2[i])
	}
	fmt.Println()
	fmt.Println()

	fmt.Println("diagraph G {")
	printBSTDOT(root2)
	fmt.Println("}")

	fmt.Printf("\nFind 88 in BST - Result: %t\n\n", searchBST(root2, 88))

	fmt.Printf("In order sorted:\n")

	inOrderBST(root2)
	fmt.Println()

}
