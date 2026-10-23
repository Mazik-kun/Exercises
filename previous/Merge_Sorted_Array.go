package main

import (
	"fmt"
	"slices"
)

func main() {
    nums1:= []int{0}
    nums2:=[]int{1}

    mergeSort(nums1,nums2)
    fmt.Println(nums1)
}
func mergeSort(nums1 []int, nums2 []int){
    n:= len(nums1) - len(nums2)
    nums1=append(nums1[:n], nums2...)
    slices.Sort(nums1)
}
