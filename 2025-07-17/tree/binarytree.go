package tree

type TreeNode struct {
  Value int
  Left *TreeNode
  Right *TreeNode
}

func PreOrder(root *TreeNode) []int {
  if root == nil {
    return nil
  }
  
  result := []int{root.Value}
  result = append(result, PreOrder(root.Left)...)
  result = append(result, PreOrder(root.Right)...)

  return result
}

func InOrder(root *TreeNode) []int {
  if root == nil {
    return nil
  }

  result := InOrder(root.Left)
  result = append(result, root.Value)
  result = append(result, InOrder(root.Right)...)

  return result
}

func PostOrder(root *TreeNode) []int {
  if root == nil {
    return nil
  }

  result := PostOrder(root.Left)
  result = append(result, PostOrder(root.Right)...)
  result = append(result, root.Value)

  return result
}
