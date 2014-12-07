package main
import (
    "fmt"
)

/* func printArr(w [4]string) { */
/*     for _,word := range w { */
/*         fmt.Printf("%s ",word) */
/*     } */
/*     fmt.Printf("\n") */
/*     w[3] = "blue" */
/* } */

/* slices */

func printArr(w []string) {
    for _,word := range w {
        fmt.Printf("%s ",word)
    }
    fmt.Printf("\n")
    w[2] = "blue"
}
func main (){
    /* words := [...]string{ "the","quick","fox" } */
    /* words := [4]string{ "the","quick","fox" } */
    words := []string{ "the","quick","brown","fox","jump","over","the","lazy","dog"}
    /* printArr(words); */
    /* printArr(words); */
    /* printArr(words); */
    /* fmt.Printf("%s\n",words[1]) */
    /* fmt.Printf("%d \n",len(words)) */
    /* fmt.Printf("%d \n",len(words)) */
    printArr(words);
    printArr(words);

}

