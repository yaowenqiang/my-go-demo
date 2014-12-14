package main
import (
    "fmt"
    "os"
)
func printer(msg string) error {
    if msg == "" {
        return fmt.Errorf("Unwilling to print an empty string");

    }
    _,err := fmt.Printf("%s\n",msg)
    return err

}
func main () {

    /* fmt.Printf("Hello world"); */
    /* fmt.Printf("Hello world"); */
    /* if  error := printer("Hello World!"); error != nil { */
    if  error := printer(""); error != nil {
        
        fmt.Printf("Printer say :failed: %s\n",error);
        os.Exit(1);
    }
    printer("");

}
