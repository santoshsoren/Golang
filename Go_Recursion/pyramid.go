package main

import (
	"fmt"
)

func print_space(space int){
	if(space==0){
		return;
	}
	fmt.Print(" ")
	print_space(space-1)
}

func print_star(n int){
	if(n==0){
		return;
	}
	fmt.Print("* ")
	print_star(n-1)
}

func star(n int,i int){
	if (n == 0){
		return;
	}
	print_space(n-1)
	print_star(i-n+1)
	fmt.Println("")
	star(n-1,i)        
} 

func main(){
   var i int = 5
   star(i,i)
   
}