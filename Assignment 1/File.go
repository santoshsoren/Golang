package main
 
import (
	"fmt"
	"os"
	"log"
)

func main(){
	
    f, err := os.Create("test.txt")

    if err != nil {
        log.Fatal(err)
    }

    defer f.Close()

    _, err2 := f.WriteString("Santosh\n")

    if err2 != nil {
        log.Fatal(err2)
    }

    fmt.Println("File Created...!!!")
}