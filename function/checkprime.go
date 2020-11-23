package main

import (
	"fmt"
)

func check(n int) int{
	for i := 2; i<=n/1; i++{
		if(n%i==0){
			return 0;
		}else{
			i++
		}
	}
	return 1
}

func main(){
   var num int = 5
   var prime int
   prime = check(num)
   if(prime == 1){
	fmt.Println("Prime number : ", num)
   }else{
	fmt.Println("not prime : ", num)
   }  
}