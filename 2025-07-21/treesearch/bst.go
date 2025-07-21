package treesearch

type BSTNode struct {
	Value int
	Left  *BSTNode
	Right *BSTNode
}

type BST struct {
	Root *BSTNode
}

func (t *BST) Insert(value int) {
	t.Root = insertNode(t.Root, value)
}

func insertNode(node *BSTNode, value int) *BSTNode {
	if node == nil {
		return &BSTNode{Value: value}
	}
	if value < node.Value {
		node.Left = insertNode(node.Left, value)
	} else if value > node.Value {
		node.Right = insertNode(node.Right, value)
	}
	return node
}

func (t *BST) Search(value int) bool {
	return searchNode(t.Root, value)
}

func searchNode(node *BSTNode, value int) bool {
	if node == nil {
		return false
	}
	if value == node.Value {
		return true
	} else if value < node.Value {
		return searchNode(node.Left, value)
	}
	return searchNode(node.Right, value)
}

func (t *BST) InOrderTraversal() []int {
	var result []int
	inOrder(t.Root, &result)
	return result
}

func inOrder(node *BSTNode, result *[]int) {
	if node == nil {
		return
	}
	inOrder(node.Left, result)
	*result = append(*result, node.Value)
	inOrder(node.Right, result)
}
