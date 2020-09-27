/**
2. Add Two Numbers
https://leetcode.com/problems/add-two-numbers/
You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

Example:

Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
Output: 7 -> 0 -> 8
Explanation: 342 + 465 = 807.

 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
*/

package main

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0

	head := &ListNode{}
	curr := head
	for l1 != nil && l2 != nil {
		curr.Next = &ListNode{}
		curr.Next.Val = (l1.Val + l2.Val + carry) % 10
		carry = (l1.Val + l2.Val + carry) / 10
		l1 = l1.Next
		l2 = l2.Next
		curr = curr.Next
	}

	for l1 != nil {
		curr.Next = &ListNode{}
		curr.Next.Val = (l1.Val + carry) % 10
		carry = (l1.Val + carry) / 10
		l1 = l1.Next
		curr = curr.Next
	}

	for l2 != nil {
		curr.Next = &ListNode{}
		curr.Next.Val = (l2.Val + carry) % 10
		carry = (l2.Val + carry) / 10
		l2 = l2.Next
		curr = curr.Next
	}

	if carry != 0 {
		curr.Next = &ListNode{}
		curr.Next.Val = carry
	}

	return head.Next
}
