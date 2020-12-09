package main

import(
	"fmt"
	"math/rand"
)

func randint(ch chan int){
	for i:=0;i<5;i++{
		fmt.Println(<-ch)
	}
}

func main(){
	ch := make(chan int)
	go randint(ch)
	ch <- rand.Intn(100)
}

