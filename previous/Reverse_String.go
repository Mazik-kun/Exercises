package main

import "fmt"


func main() {
    str:= []byte{'h','e','l','l','o',' ', 'w', 'o', 'r', 'l', 'd'}
    reversy(str)
}
func reversy(str []byte){
    for i:=0; i<(len(str)-1) /2; i++{
        str[i],str[len(str)-i-1] = str[len(str)-i-1],str[i]
    }
    fmt.Print(str)
}
