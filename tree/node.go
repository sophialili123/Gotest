package main

import "fmt"

type treeNode struct {
	value int
	left,rignt *treeNode
}

func createNode(value int) *treeNode {
	return &treeNode{value:value}//局部变量的地址
}
/**
	为结构体创建方法
	强类型限制
 */
func(node treeNode) print(){
	fmt.Println(node.value)
}

func (node *treeNode)setValue(value int)  {
	if node == nil{
		fmt.Println("Node is Nil")
		return
	}
	node.value = value
}

func (node *treeNode) traverse()  {
	if node == nil{
		return
	}
	node.left.traverse()
	node.print()
	node.rignt.traverse()
}

func main() {
	var root treeNode
	root = treeNode{value:3}
	root.left = &treeNode{}
	root.rignt = &treeNode{5,nil,nil}
	root.rignt.left = new(treeNode)
	root.left.rignt = createNode(2)//接收了局部地址的

	//fmt.Println(root)
	//
	//nodes := []treeNode{
	//	{value:3},{},{6,nil,&root},
	//}
	//fmt.Println(nodes)

	root.print();
	root.rignt.left.setValue(9)
	root.rignt.left.print()
	fmt.Println()

	root.print()
	root.setValue(100)
	root.print()

	fmt.Println("=========")
	pRoot := &root
	pRoot.print()
	pRoot.setValue(200)
	pRoot.print()

	fmt.Println("=========")
	var sRoot *treeNode
	sRoot.setValue(300)
	sRoot = &root
	sRoot.setValue(400)
	sRoot.print()

	fmt.Println("=========")
	root.traverse()
}
