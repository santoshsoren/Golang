package main

import(
	"fmt"
)

func main() {
	end := 25
	for i := 1;i < end;i++{
		if(i%2==0){
			fmt.Println("This is even number : ", i)
		}else{
			fmt.Println("This is odd number : ", i)
		}
	}
}