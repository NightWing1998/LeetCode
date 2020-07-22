// Link - https://leetcode.com/problems/binary-tree-zigzag-level-order-traversal.go

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func zigzagLevelOrder(root *TreeNode) [][]int {
    order := make([][]int,0);
    leftToRight(root,0,&order);
    return order;
}

func leftToRight(root *TreeNode,level int,order *[][]int){
    if root == nil {
        return;
    }
    if len(*order) <= level {
        *order = append(*order,make([]int,0,1));
    }
    (*order)[level] = append((*order)[level],root.Val);
    if root.Left != nil {rightToLeft(root.Left,level+1,order);}
    if root.Right != nil {rightToLeft(root.Right,level+1,order);}
}

func rightToLeft(root *TreeNode,level int,order *[][]int){
    if root == nil {
        return;
    }
    if len(*order) <= level {
        *order = append(*order,make([]int,0,1));
    }
    (*order)[level] = append([]int{root.Val},(*order)[level]...);
    if root.Left != nil {leftToRight(root.Left,level+1,order)}
    if root.Right != nil {leftToRight(root.Right,level+1,order);}
}

// Time - 0ms(beats 100%)
// Memory - 2.8MB(beats 15.38%)

