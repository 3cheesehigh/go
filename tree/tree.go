package tree

type Node struct {
	Content     int
	Left, Right *Node
}

func Insert(tree *Node, n int) *Node {
	if tree == nil {
		return &Node{n, nil, nil}
	}
	if n < tree.Content {
		tree.Left = Insert(tree.Left, n)
		return tree
	}
	tree.Right = Insert(tree.Right, n)
	return tree
}
