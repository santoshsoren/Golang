package main

import(
	"fmt"
)

func main() {
	var sum int = 10
	for i := 1;i < 10;i++{
		sum = sum + i
		fmt.Println(sum)
	}
}