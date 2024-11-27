/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func preorder(root *Node) []int {
	if root == nil {
		return nil
	}

	nodesStack := []*Node{root}
	var traversal []int

	for len(nodesStack) > 0 {
		lastStackNodeIndex := len(nodesStack) - 1
		nodeCurrent := nodesStack[lastStackNodeIndex]
		nodesStack = nodesStack[:lastStackNodeIndex]

		traversal = append(traversal, nodeCurrent.Val)

		if len(nodeCurrent.Children) == 0 {
			continue
		}

		for i := len(nodeCurrent.Children) - 1; i >= 0; i-- {
			nodesStack = append(nodesStack, nodeCurrent.Children[i])
		}
	}

	return traversal
}

