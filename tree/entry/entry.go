package main

import (
	"tree"
	"fmt"
)

type myTreeNode struct {
	node *tree.Node
}

func (myNode *myTreeNode) postOrder() {
	if myNode == nil || myNode.node == nil {
		return
	}

	left := myTreeNode{myNode.node.Left}
	left.postOrder()
	right := myTreeNode{myNode.node.Right}
	right.postOrder()
	myNode.node.Print()
}


func main() {
	var root tree.Node

	root = tree.Node{Value: 3}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5, nil, nil}
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreateNode(2)
	root.Right.Left.SetValue(4)

	fmt.Println("In-order traversal: ")
	root.Traverse()
	fmt.Println("------------")
	myRoot := myTreeNode{&root}
	myRoot.postOrder()
	fmt.Println()

	//计数器的作用
	root.Traverse()
	nodeCount :=0
	root.TraverseFunc(func(node *tree.Node) {
		nodeCount ++
	})
	fmt.Println("Node count:",nodeCount)
}
