package main

import (
    "fmt"
    "os"
)

func main (){
    /* var numberBytes int */
    /* var theError error */
    /* if numberBytes,theError := fmt.Printf("Hello world!\n");theError != nil { */
    /* if numberBytes,theError = fmt.Printf("Hello world!\n");theError != nil { */
    if _,theError := fmt.Printf("Hello world!\n");theError != nil {
        os.Exit(1)
    }

}

