package tree

func Inorder(rootNode *Node, ch chan int) {
	if rootNode == nil {
		return
	}

	Inorder(rootNode.Left, ch)
	ch <- rootNode.Data
	Inorder(rootNode.Right, ch)
}

func Preorder(rootNode *Node, ch chan int) {
	if rootNode == nil {
		return
	}

	ch <- rootNode.Data
	Preorder(rootNode.Left, ch)
	Preorder(rootNode.Right, ch)
}

func Postorder(rootNode *Node, ch chan int) {
	if rootNode == nil {
		return
	}

	Postorder(rootNode.Left, ch)
	Postorder(rootNode.Right, ch)
	ch <- rootNode.Data
}
