package main

import (
	"fmt"
)

func print(n int){
	if(n==0){
		return;
	}
	fmt.Print("*")
	print(n-1)
}
func star(n int,i int){
	if (n == 0){
		return;
	}
	print(i)
	fmt.Println()
	star(n-1,i+1)        
} 

func main(){
   var i int = 5
   star(i,1)
   
}