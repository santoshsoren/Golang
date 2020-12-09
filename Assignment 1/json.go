package main
 
import (
	"encoding/json"
	"fmt"
)

type Book struct{
	Title   string  `json:"title"`
 	Author  string  `json:"author"`
}

func main(){
	book := Book{Title:"Learning json", Author:"xyz"}
	
	byteArray, error := json.MarshalIndent(book, " ", "  ")
	if(error != nil){
		fmt.Println(error)
	}
	fmt.Println(string(byteArray))
}