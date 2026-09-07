package main

import (
	"fmt"
)
func main(){
	list1 := NewLinkedList([]int{1,2,4})
	list2 := NewLinkedList([]int{1,3,4})

	Merge(list1,list2).Print()
}


type Node struct {
	Val int
	Next *Node
}
type LinkedList struct {
	Head *Node
}
func NewLinkedList(values []int) * LinkedList{
	list := &LinkedList{}
	for _, val := range values {
		list.Append(val)
	}
	return list
}
func (l * LinkedList) Append(val int){
	newNode := &Node{Val: val}
	if l.Head == nil{
		l.Head = newNode
		return
	}
	current := l.Head
	for current.Next != nil{
		current = current.Next
	}
	current.Next = newNode
}
func (l * LinkedList) Print(){
	current := l.Head 
	for current != nil{
		fmt.Printf("%d", current.Val)
		current = current.Next
	}
	fmt.Println()
}
func Merge(list1, list2  *LinkedList) *LinkedList{
	slice := []int{}
	
	for !(list1.Head.Val == 51 && list2.Head.Val == 51){
		if list1.Head.Val < list2.Head.Val {
			slice = append(slice, list1.Head.Val)
			if list1.Head.Next == nil{
				list1.Head.Val = 51;
			} else{
				list1.Head = list1.Head.Next
				// fmt.Printf("slice: %v\n", slice)
			}
			
		}else{
			slice = append(slice, list2.Head.Val)
			if list2.Head.Next == nil{
				list2.Head.Val = 51
			}else{	
				list2.Head = list2.Head.Next
				// fmt.Printf("slice: %v\n", slice)
			}
		}
	}	
	// fmt.Printf("slice: %v\n", slice)
	newLink := NewLinkedList(slice)
return newLink
}
