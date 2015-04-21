package main

import "fmt"

import . "github.com/3cheesehigh/go/tree"

/*
Zwei Bäume heißein äquivalent, wenn sie gleichviele Knoten haben und deren Inhalte ( in Preorder -Reihenfolge ) paarweise übereinstimmen.
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
	Insert(&tree1, 5)
	Insert(&tree2, 5)

	fmt.Println(equalTrees(&tree1, &tree2))
	//mt.Println(equalTrees(tree1, tree3))
	//mt.Println(equalTrees(tree1, tree4))
	//mt.Println(equalTrees(tree1, tree5))
}
