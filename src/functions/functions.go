package main
import (
    "fmt"
    /* "os" */
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

/* func printer(message string) error { */
/*     defer fmt.Printf("More\n") */
/*     defer fmt.Printf("\n") */
/*     _,error := fmt.Printf("%s",message) */
/*     return error */
/* } */

/* operate file */
/* func printer(message string) error { */
/*     fmt.Printf(message) */
/*     f,err := os.Create("hellorword.txt") */
/*     if err != nil { */
/*         return err */
/*     } */
/*     defer f.Close() */
/*     //TODO 文件写入失败 */
/*     /1* f.Write(message) *1/ */
/*     return err */
/* } */


/* return named value */
func printer(message string) (e error) {
    _,e = fmt.Printf("%s\n",message)
    return
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
