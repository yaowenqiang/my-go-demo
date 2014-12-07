package main
import (
    "fmt"
)

/* func printer(message1,message2 string,string) { */
/* func printer(message1,message2 string,repeat int) error { */
/*     for repeat>0 { */
/*     fmt.Printf("%s %s\n",message1,message2) */
/*     repeat -- */
/*     } */
/* } */

func printer(message string) error {
    _,error := fmt.Printf("%s\n",message)
    return error
}
func main (){
    /* printer("Hello") */
    printer("Hello")

}
