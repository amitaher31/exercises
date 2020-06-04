/*
https://leetcode.com/problems/design-hashmap/
706. Design HashMap

Design a HashMap without using any built-in hash table libraries.

To be specific, your design should include these functions:

put(key, value) : Insert a (key, value) pair into the HashMap. If the value already exists in the HashMap, update the value.
get(key): Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key.
remove(key) : Remove the mapping for the value key if this map contains the mapping for the key.

Example:

MyHashMap hashMap = new MyHashMap();
hashMap.put(1, 1);
hashMap.put(2, 2);
hashMap.get(1);            // returns 1
hashMap.get(3);            // returns -1 (not found)
hashMap.put(2, 1);          // update the existing value
hashMap.get(2);            // returns 1
hashMap.remove(2);          // remove the mapping for 2
hashMap.get(2);            // returns -1 (not found)

Note:

All keys and values will be in the range of [0, 1000000].
The number of operations will be in the range of [1, 10000].
Please do not use the built-in HashMap library.
*/
package main

import "fmt"

func main() {
	mymap := Constructor()
	fmt.Println(mymap)
	mymap.Put(1, 1)
	fmt.Println(mymap.Get(1))
	mymap.Put(1, 2)
	fmt.Println(mymap.Get(1))
	mymap.Put(3, 2)
	fmt.Println(mymap.Get(3))
	mymap.Remove(3)
	fmt.Println(mymap.Get(3))
}

const capacity = 5

type Node struct {
	Key   int
	Value int
	Next  *Node
}

//here key_space is an [ ] of memory addresses of Node which can be chained since Node is a linked list
type MyHashMap struct {
	Key_Space [capacity]*Node
}

/** Initialize your data structure here. */
func Constructor() MyHashMap {
	m := MyHashMap{}
	return m
}

/** value will always be non-negative. */
func (this *MyHashMap) Put(key int, value int) {
	// find hash of the key and check if its present in MyHashMap i-e this
	hash := key % capacity
	curr := this.Key_Space[hash]

	// new node to be used if existing not found
	newNode := Node{
		Key:   key,
		Value: value,
	}
	//if key_space if empty (1st node) add it to the key_space []
	if curr == nil {
		this.Key_Space[hash] = &newNode
		return
	}

	//if a existing node is present
	var prev *Node
	for curr != nil {
		//if key same as existing key - update the current value
		if curr.Key == key {
			curr.Value = value
			return
		}
		prev = curr
		curr = curr.Next
	}

	// if new key - add it to the end
	prev.Next = &newNode

}

/** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
func (this *MyHashMap) Get(key int) int {
	hash := key % capacity
	curr := this.Key_Space[hash]

	if curr == nil {
		return -1
	}

	for curr != nil {
		if curr.Key == key {
			return curr.Value
		}
		curr = curr.Next
	}
	return -1
}

/** Removes the mapping of the specified value key if this map contains a mapping for the key */
func (this *MyHashMap) Remove(key int) {
	hash := key % capacity
	curr := this.Key_Space[hash]

	if curr == nil {
		return
	}

	if curr.Key == key {
		this.Key_Space[hash] = curr.Next
		return
	}

	var prev *Node
	for curr.Next != nil {
		prev = curr
		curr = curr.Next
		if curr.Key == key {
			prev.Next = curr.Next
		}

	}
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */
