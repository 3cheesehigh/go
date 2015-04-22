package main

import "fmt"

import . "github.com/3cheesehigh/go/tree"

/*
Zwei Bäume heißen äquivalent, wenn sie gleich viele Knoten
haben und deren Inhalte ( in Preorder -Reihenfolge ) paarweise übereinstimmen.
*/

func equalTrees(tree1, tree2 *Node) bool {
	ch1, ch2 := travRoutines(tree1), travRoutines(tree2)
	for {
		cont1, ok1 := <-ch1
		cont2, ok2 := <-ch2
		if !ok1 || !ok2 {
			return ok1 == ok2
		}
		if cont1 != cont2 {
			break
		}
	}
	return false

}
func Traverse(tree *Node, ch chan int) {
	if tree == nil {
		return
	}
	Traverse(tree.Left, ch)
	ch <- tree.Content
	Traverse(tree.Right, ch)
}

func travRoutines(tree *Node) chan int {
	ch := make(chan int)
	go func() {
		Traverse(tree, ch)
		close(ch)
	}()
	return ch
}

func main() {
	tree1 := Node{1, nil, nil}
	tree2 := Node{1, nil, nil}
	tree3 := Node{1, nil, nil}
	tree4 := Node{1, nil, nil}
	xs := []int{4, 2, 3, 1, 5, 8, 2}
	ys := []int{2, 3, 7, 4, 8, 2, 3, 1, 8, 9}

	for i, x := range xs {

		Insert(&tree1, x)
		Insert(&tree2, x)     // Gleicher Baum
		Insert(&tree3, ys[i]) // Ungleich aber gleiche Anzahl
	}
	for _, y := range ys {
		Insert(&tree4, y) // Ungleich Bäume
	}
	fmt.Println(equalTrees(&tree1, &tree2))
	fmt.Println(equalTrees(&tree1, &tree3))
	fmt.Println(equalTrees(&tree1, &tree4))
}
