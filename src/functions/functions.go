package main
import (
    "fmt"
)

/* func printer(message1,message2 string,string) { */
func printer(message1,message2 string) {
    fmt.Printf("%s %s\n",message1,message2)
}
func main (){
    printer("Hello"," world")

}
