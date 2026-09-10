package main

import (
	"fmt"
)
type ListNode struct {
    Val int
    Next *ListNode
}
func main() {
    vals := []int{7,6,5,4,3,2,1,1}
    head := sliceToList(vals)
    head.Next.Next.Next.Next.Next.Next = head.Next
    print(IsCycle(head))
    // printList(head)
}
func sliceToList(vals []int) *ListNode{
    head:= &ListNode{Val: vals[0]}
    cur:= head
    for i:=1; i<len(vals); i++{
        cur.Next = &ListNode{Val:vals[i]}
        cur = cur.Next
    }
    return head
}

func printList(head *ListNode) {
    cur := head
    for cur != nil {
        fmt.Printf("%d -> ", cur.Val)
        cur = cur.Next
    }
    fmt.Println("nil")
}

func IsCycle(head *ListNode)bool{
    fast:=head
    slow:=head
    for i:=0;fast != nil;i++{
        if i % 2 == 1{
            slow = slow.Next
        }
        fast = fast.Next
        
        if slow == fast{
            return true
        }
        
    
    }
return false}
