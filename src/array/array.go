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
    /* words := []string{ "the","quick","brown","fox","jump","over","the","lazy","dog"} */
    /* printArr(words); */
    /* printArr(words); */
    /* printArr(words); */
    /* fmt.Printf("%s\n",words[1]) */
    /* fmt.Printf("%d \n",len(words)) */
    /* fmt.Printf("%d \n",len(words)) */
    /* printArr(words); */
    /* printArr(words); */
    /* printArr(words[1:7]); */
    /* printArr(words[5:len(words)]); */
    /* printArr(words[5:]); */
    /* printArr(words[:9]); */
    /* words := make([]string,4) */
    words := make([]string,0,4)
    fmt.Printf("%d,%d\n",len(words),cap(words))
    words = append(words,"The")
    words = append(words,"Quick")
    words = append(words,"Brown")
    words = append(words,"Fox")
    words = append(words,"Fox")
    words = append(words,"Fox")
    words = append(words,"Fox")
    words = append(words,"Fox")
    words = append(words,"Fox")
    /* words[0] = "The" */
    /* words[1] = "Quick" */
    /* words[2] = "Brown" */
    /* words[3] = "Fox" */
    printArr(words)
    newWords := make([]string,4)
    copy(newWords,words)
    words[3] = "Blue"
    printArr(newWords)

}

