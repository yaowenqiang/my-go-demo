package main
import (
    "fmt"
)
func printArr(w [4]string) {
    for _,word := range w {
        fmt.Printf("%s ",word)
    }
    fmt.Printf("\n")
}
func main (){
    /* words := [...]string{ "the","quick","fox" } */
    words := [4]string{ "the","quick","fox" }
    printArr(words);
    /* fmt.Printf("%s\n",words[1]) */

}

