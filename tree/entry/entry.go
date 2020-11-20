package main

import (
	"fmt"
	"github.com/839891627/tree"
)

type myTreeNode struct {
	*tree.Node // Embedding 内嵌
}

func (myNode *myTreeNode) postOrder() {
	if myNode == nil || myNode.Node == nil {
		return
	}

	left := myTreeNode{myNode.Left}
	right := myTreeNode{myNode.Right}

	left.postOrder()
	right.postOrder()

	myNode.Print()
}

func (myNode *myTreeNode) Traverse() {
	fmt.Println("This method is shadowed.")
}

func main() {
	//var root tree.Node
	//fmt.Println(root)
	root := myTreeNode{&tree.Node{Value: 3}}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{Value: 5}
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreateNode(2)
	//root.print()
	root.Right.Left.SetValue(4)
	root.Traverse()
	root.Node.Traverse()

	fmt.Println()
	root.postOrder()
	fmt.Println()

	//var baseRoot *tree.Node

}
