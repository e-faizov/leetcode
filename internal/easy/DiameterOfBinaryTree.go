package easy

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	_, res := getMaxBranchAndPath(root)
	return res
}

func getMaxBranchAndPath(node *TreeNode) (int, int) {
	var leftMaxPath, rightMaxPath, nodePath, leftBranch, rightBranch int
	if node.Left != nil {
		leftBranch, leftMaxPath = getMaxBranchAndPath(node.Left)
		leftBranch++
	}

	if node.Right != nil {
		rightBranch, rightMaxPath = getMaxBranchAndPath(node.Right)
		rightBranch++
	}

	nodePath = leftBranch + rightBranch
	return max(rightBranch, leftBranch), max(nodePath, leftMaxPath, rightMaxPath)
}
