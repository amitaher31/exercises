/*
206. Reverse Linked List - https://leetcode.com/problems/reverse-linked-list/
Reverse a singly linked list.
Example:

Input: 1->2->3->4->5->NULL
Output: 5->4->3->2->1->NULL
*/

package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//iterative
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		tmp := curr.Next
		curr.Next = prev
		prev = curr
		curr = tmp
	}
	return prev
}

/* recursive
Eg 1 -> 2 -> 3
 - use callstack to set reverse the pointers
 - till we reach last node (i-e curr.Next == nil) call reverseList(curr.Next)
    callstack:
        - reverseList(1) head=1
        - reverseList(2) head=1
 - once we reach last node return head i-e (head = curr) as the new head
    callstack:
        - reverseList(3) head=1 => return head=3
 - back to the call stack where curr points to n-1 node:
  - set curr.next.next = curr as (i) & curr.next = nil (ii)
    callstack:
        - reverseList(2) head=3 1->2->3 => 1->2->3->2 (i) => 1->2 3->2 (ii)
        - reverseList(1) head=3 1->2 3->2 => 1->2 3->2->1 => 1 2 3->2->1
        - return head=3 ie head->3->2->1
*/
func reverseListRecursive(head *ListNode) *ListNode {
	fmt.Println("h = ", head)
	//empty list
	if head == nil {
		return nil
	}

	//exit condition & final return point
	if head.Next == nil {
		fmt.Println("returning head ", head)
		return head
	}

	r := reverseListRecursive(head.Next)
	fmt.Println(" - bacck ", r, head)
	head.Next.Next = head
	head.Next = nil
	fmt.Println(" - 2 bacck ", r, head)
	return r
}
