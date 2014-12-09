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
    defer f.Close()
    b := make([]byte,100)

    n,err := f.Read(b)
    /* fmt.Printf("%d: % x\n",n,b) */
    /* fmt.Printf("%d: % s\n",n,b) */
    fmt.Printf("%d: % c\n",n,b)
    /* stringVersion := string(b) */
    fmt.Printf("%d: % s\n",n,stringVersion)
}
