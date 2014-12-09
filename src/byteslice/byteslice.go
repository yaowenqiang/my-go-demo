package main
import (
    "fmt"
    "os"
)


func main (){
    f,err := os.Open("test.txt")
    if err != nil {
        fmt.Printf("%s\n",err)
        os.Exit(1)
    }
    b := make([]byte,100)

    n,err := f.Read(b)
    /* fmt.Printf("%d: % x\n",n,b) */
    /* fmt.Printf("%d: % s\n",n,b) */
    /* fmt.Printf("%d: % c\n",n,b) */
    stringVersion := string(b)
    fmt.Printf("%d: % s\n",n,stringVersion)
    someString := "some string"
    /* f.Write(someString) */
    _,err = f.Write([]byte(someString))
    fmt.Printf("%s\n",err)
    defer f.Close()
}
