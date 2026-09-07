package main

import (
	// "fmt"
)

func main(){
	days:= []int{7,1,5,3,6,4}
	println(BuyNSell(days))
}

func BuyNSell(days[]int) int{
	BestDeal := 0;
	for i:= 0;  i< len(days)-1; i++{
		for j:= i+1; j< len(days);j++{
			if days[j]-days[i] > BestDeal{
				BestDeal = days[j]-days[i];
			}
		}
	}
return BestDeal}
