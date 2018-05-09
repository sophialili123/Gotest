package tree

import "fmt"

type Node struct {
	Value       int
	Left, Right *Node
}

func (node Node) Print() {
	fmt.Print(node.Value, " ")
}

func (node *Node) SetValue(value int) {
	if node == nil {
		fmt.Println("Setting Value to nil " +
			"node. Ignored.")
		return
	}
	node.Value = value
}

func CreateNode(value int) *Node {
	return &Node{Value: value}
}

func (node *Node)Traverse()  {
	if node == nil{
		return
	}

	node.Left.Traverse()
	node.Print()
	node.Right.Traverse()

}

//func main() {
//	var root treeNode
//	root = treeNode{value:3}
//	root.left = &treeNode{}
//	root.rignt = &treeNode{5,nil,nil}
//	root.rignt.left = new(treeNode)
//	root.left.rignt = createNode(2)//接收了局部地址的
//
//	//fmt.Println(root)
//	//
//	//nodes := []treeNode{
//	//	{value:3},{},{6,nil,&root},
//	//}
//	//fmt.Println(nodes)
//
//	root.print();
//	root.rignt.left.setValue(9)
//	root.rignt.left.print()
//	fmt.Println()
//
//	root.print()
//	root.setValue(100)
//	root.print()
//
//	fmt.Println("=========")
//	pRoot := &root
//	pRoot.print()
//	pRoot.setValue(200)
//	pRoot.print()
//
//	fmt.Println("=========")
//	var sRoot *treeNode
//	sRoot.setValue(300)
//	sRoot = &root
//	sRoot.setValue(400)
//	sRoot.print()
//
//	fmt.Println("=========")
//	root.traverse()
//}
