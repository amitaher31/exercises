/*
https://leetcode.com/problems/find-all-the-lonely-nodes/
n a binary tree, a lonely node is a node that is the only child of its parent node. The root of the tree is not lonely because it does not have a parent node.

Given the root of a binary tree, return an array containing the values of all lonely nodes in the tree. Return the list in any order.



Example 1:


Input: root = [1,2,3,null,4]
Output: [4]
Explanation: Light blue node is the only lonely node.
Node 1 is the root and is not lonely.
Nodes 2 and 3 have the same parent and are not lonely.
Example 2:


Input: root = [7,1,4,6,null,5,3,null,null,null,null,null,2]
Output: [6,2]
Explanation: Light blue nodes are lonely nodes.
Please remember that order doesn't matter, [2,6] is also an acceptable answer.
Example 3:



Input: root = [11,99,88,77,null,null,66,55,null,null,44,33,null,null,22]
Output: [77,55,33,66,44,22]
Explanation: Nodes 99 and 88 share the same parent. Node 11 is the root.
All other nodes are lonely.
Example 4:

Input: root = [197]
Output: []
Example 5:

Input: root = [31,null,78,null,28]
Output: [78,28]

*/

package main

import (
	"fmt"
	"io"
	"os"
)

//--------------------

func main() {
	tree := &BinaryTree{}
	tree.insert(100).
		insert(-20).
		insert(-50).
		insert(-15).
		insert(-60).
		insert(50).
		insert(60).
		insert(55).
		insert(85).
		insert(15).
		insert(5).
		insert(-10)

	//[1,2,3,null,4]
	// t1 := []int{3, 2, 4, 3, 4}
	// t1 := []int{1, 1, 1, 1, 1, 1, 1}

	// [1,2,3,4,5,6,7]
	// t1 := []int{1, 2, 3, 4, 5, 6, 7}
	// for _, v := range t1 {
	// 	tree.insert(v)
	// }
	// tree.insert(3).
	// 	insert(2).
	// 	insert(4).
	// 	insert(3).
	// 	insert(4)
	print(os.Stdout, tree.root, 0, 'M')

	fmt.Println("Right view = ", rightSideView(tree.root))
}

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//Custom Binary tree to support Leet Code problems
//Ref: https://www.golangprograms.com/data-structure-and-algorithms/golang-program-to-implement-binary-tree.html
type BinaryTree struct {
	root *TreeNode
}

func (t *BinaryTree) insert(data int) *BinaryTree {
	if t.root == nil {
		t.root = &TreeNode{Val: data, Left: nil, Right: nil}
	} else {
		t.root.insert(data)
	}
	return t
}

func (n *TreeNode) insert(data int) {
	if n == nil {
		return
	} else if data <= n.Val {
		if n.Left == nil {
			//add node
			n.Left = &TreeNode{Val: data, Left: nil, Right: nil}
		} else {
			//add node to the left sub tree below n
			n.Left.insert(data)
		}
	} else {
		if n.Right == nil {
			//add node
			n.Right = &TreeNode{Val: data, Left: nil, Right: nil}
		} else {
			//add node to the right sub tree
			n.Right.insert(data)
		}
	}
}

func print(w io.Writer, node *TreeNode, ns int, ch rune) {
	if node == nil {
		return
	}

	for i := 0; i < ns; i++ {
		fmt.Fprint(w, " ")
	}
	fmt.Fprintf(w, "%c:%v\n", ch, node.Val)
	print(w, node.Left, ns+2, 'L')
	print(w, node.Right, ns+2, 'R')
}

//---------------------------

func rightSideView(root *TreeNode) []int {
	r := []int{}
	l := []int{}
	if root == nil {
		return r
	}
	r1 := [][]int{}

	levelNodes := []*TreeNode{root}

	for len(levelNodes) > 0 {
		r = append(r, levelNodes[len(levelNodes)-1].Val)
		l = append(l, levelNodes[0].Val)
		l1 := []int{}
		for _, v := range levelNodes {
			l1 = append(l1, v.Val)
		}
		r1 = append(r1, l1)

		fmt.Println("before - ", levelNodes)
		levelNodes = nextLevelNodes(levelNodes)
		fmt.Println("after - ", levelNodes)
	}

	fmt.Println("r1 = ", r1)
	fmt.Println("r = ", r)
	fmt.Println("l = ", l)

	return r
}

func nextLevelNodes(curr []*TreeNode) []*TreeNode {
	nextLevel := []*TreeNode{}

	for _, n := range curr {
		if n.Left != nil {
			nextLevel = append(nextLevel, n.Left)
		}
		if n.Right != nil {
			nextLevel = append(nextLevel, n.Right)
		}
	}
	return nextLevel
}

//---------------------------

//  func getLonelyNodes(root *TreeNode) []int {
//     lonely = []int{}
//     if root == nil {
//         return lonely
//     }

//     if root.Left == nil && root.Right != nil {
//         lonely = append(lonely, root.Right.Val)
//     }

//     if root.Right == nil && root.Left != nil {
//         lonely = append(lonely, root.Left.Val)
//     }

//     visit(root.Left)
//     visit(root.Right)
//     return lonely
// }
// var lonely []int

// func visit(curr *TreeNode) {
//     if curr == nil {
//         return
//     }

//     if curr.Left == nil && curr.Right != nil {
//         lonely = append(lonely, curr.Right.Val)
//     }

//     if curr.Right == nil && curr.Left != nil {
//         lonely = append(lonely, curr.Left.Val)
//     }

//     visit(curr.Left)
//     visit(curr.Right)
// }
