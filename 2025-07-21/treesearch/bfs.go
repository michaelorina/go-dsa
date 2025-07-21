package treesearch

type BFSNode struct {
	Value int
	Left  *BFSNode
	Right *BFSNode
}

func BFS(root *BFSNode) []int {
	if root == nil {
		return nil
	}

	var result []int
	queue := []*BFSNode{root}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		result = append(result, current.Value)

		if current.Left != nil {
			queue = append(queue, current.Left)
		}
		if current.Right != nil {
			queue = append(queue, current.Right)
		}
	}

	return result
}
