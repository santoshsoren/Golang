package main

import (
	"fmt"
)

func area_square(side int)int{
	var area int
	area = side*side
	return area
}

func main(){
   var side int = 5
   fmt.Println(area_square(side))
   
}