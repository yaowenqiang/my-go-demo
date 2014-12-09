package main
import (
    "fmt"
    "os"
)


func main (){
    _,err := os.Open("test.txt")
    if err != nil {
        fmt.Printf("%s\n",err)
        os.Exit(1)
    }

}
