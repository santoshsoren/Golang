package main

import(
	"fmt"
	"strconv"
)

func main() {
	var i int
	i = 27
	var j float32
	j = float32(i)
	var k string
	k = strconv.Itoa(i)
	fmt.Printf("%v , %T\n", i , i)
	fmt.Printf("%v , %T\n", j , j)
	fmt.Printf("%v , %T\n", k , k)
}