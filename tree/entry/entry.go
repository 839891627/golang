package main

import (
	"fmt"
	"github.com/839891627/tree"
)

type myTreeNode struct {
	node *tree.Node
}

func (myNode *myTreeNode) postOrder() {
	if myNode == nil || myNode.node == nil {
		return
	}

	left := myTreeNode{myNode.node.Left}
	right := myTreeNode{myNode.node.Right}

	left.postOrder()
	right.postOrder()

	myNode.node.Print()
}

func main() {
	var root tree.Node
	//fmt.Println(root)
	root = tree.Node{Value: 3}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{Value: 5}
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreateNode(2)
	//root.print()
	root.Right.Left.SetValue(4)
	root.Traverse()

	fmt.Println()

	myRoot := myTreeNode{&root}
	myRoot.postOrder()
}
