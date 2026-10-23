package main

import (
	"fmt"
)
type ListNode struct {
    Val int
    Next *ListNode
}
func main() {
    vals := []int{4,3,2,1}
    head := sliceToList(vals)
    printList(head)
    ToDelete(head, 4)
    printList(head)
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
func ToDelete(head *ListNode, idx int)*ListNode{
    cur:= head;
    pntr:= head;
    for i:=0;;i++{       
        if i == idx{
            cur = cur.Next
            pntr = head
        }else{
            cur = cur.Next    
            pntr = pntr.Next
        }
        if cur == nil{
            if pntr == nil{
               *head = *head.Next
                return head.Next
            }
            pntr.Next = pntr.Next.Next
            return head
        }
    }
}
