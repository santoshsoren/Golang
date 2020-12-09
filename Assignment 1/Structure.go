package main
 
import (
	"fmt"
)

type Student struct{
	name   string
	marks  []int
	age    int
}

func (s Student) getavgmark() float32{
	sum := 0
	for _, v := range s.marks{
		sum += v
	}
	return float32(sum)/float32(len(s.marks))
}

func main(){
	s1 := Student{"Sandy",[]int{85,92,87,98},19}
	//avg = getavgmark(s1) 
	fmt.Println(s1)
	fmt.Println(s1.name)
	fmt.Println(s1.age)
	fmt.Println(s1.getavgmark())
}