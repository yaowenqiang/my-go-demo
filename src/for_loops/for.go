package main
import (
    "fmt"
    /* "os" */
)
func main (){
    /* var counter int */
    /* counter = 0 */
    /* for counter < 10 { */
    /*     fmt.Printf("Hello world!\n") */
    /*     counter ++ */
    /* } */
    /* for i :=0;i < 10 ;i++ { */
    /*     fmt.Printf("Hello world!\n") */
    /* } */
    for i,j :=0,1;i < 10 ;i,j = i+1,j+2 {
        fmt.Printf("%d Hello world!\n",j)
    }

}
