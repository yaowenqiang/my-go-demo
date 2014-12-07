package main
import (
    "fmt"
)

/* func printer(message1,message2 string,string) { */
func printer(message1,message2 string,repeat int) {
    for repeat>0 {
    fmt.Printf("%s %s\n",message1,message2)
    repeat --
    }
}
func main (){
    printer("Hello"," world",10)

}
