// Link - https://leetcode.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(inorder []int, postorder []int) *TreeNode {
    return constructTree(inorder,postorder,0,len(inorder)-1,len(postorder)-1);
}

func constructTree(inorder []int,postorder []int, start,end,index int) *TreeNode {
    if start > end {return nil;}
    root := TreeNode{};
    root.Val = postorder[index];
    
    rootIndexInorder := start;
    for root.Val != inorder[rootIndexInorder] {rootIndexInorder++}
    
    root.Right = constructTree(inorder,postorder,rootIndexInorder+1,end,index-1);
    root.Left = constructTree(inorder,postorder,start,rootIndexInorder-1,index - end + rootIndexInorder - 1);
    
    return &root;
}

// Time - 20ms(beats 88.89%)
// Memory - 9.7MB(beats 83.33%)

