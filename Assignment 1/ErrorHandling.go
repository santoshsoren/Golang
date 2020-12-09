package main
 
import (
	"fmt"
	"errors"
)

func div(n int, d int )(int, error){
	division := 0
	if(d == 0){
		return 0, errors.New("Division not possible")
	} else{
		division = n/d
		return division, nil
	}
}

func main(){
	division, error := div(5,2)
	if(error != nil){
		fmt.Println("There was an error", error)
	} else{
		fmt.Println(division)
	}
}