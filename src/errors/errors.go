package main
import (
    "fmt"
    "os"
    "errors"
)

var (
    errorEmptyString = errors.New("Unwilling to print an epty string")
)
func printer(msg string) error {
    if msg == "" {
        /* return fmt.Errorf("Unwilling to print an empty string"); */
        return errorEmptyString
    }
    _,err := fmt.Printf("%s\n",msg)
    return err

}
func main () {

    /* fmt.Printf("Hello world"); */
    /* fmt.Printf("Hello world"); */
    /* if  error := printer("Hello World!"); error != nil { */
    if  error := printer(""); error != nil {
        
        if  error == errorEmptyString {
            fmt.Printf("You tried to print an empty string!")
        } else {
            fmt.Printf("Printer say :failed: %s\n",error);
        }
        os.Exit(1);
    }
    printer("");

}
