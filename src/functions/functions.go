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

/* single return value */
/* func printer(message string) error { */
/*     _,error := fmt.Printf("%s\n",message) */
/*     return error */
/* } */

/* multiple return value */
/* func printer(message string) (string,error) { */
/*     message +="\n" */
/*     _,error := fmt.Printf("%s\n",message) */
/*     return message,error */
/* } */

/* defer */

func printer(message string) error {
    defer fmt.Printf("More\n")
    defer fmt.Printf("\n")
    _,error := fmt.Printf("%s",message)
    return error
}
func main (){
    /* printer("Hello") */
    /* appendedMessage,error :=  printer("Hello") */
    /* if error == nil { */
    /*     fmt.Printf("%q\n",appendedMessage) */
    /*     fmt.Printf("%x\n",appendedMessage) */
    /*     fmt.Printf("% x\n",appendedMessage) */

    /* } */
    printer("Hello")

}
