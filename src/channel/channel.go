package main
import (
    "fmt"
    "math/rand"
)
/* func emit(c chan string) { */
/*     words := []string{"The","Quick","Brown","Fox"} */
/*     for _,word := range words { */
        
/*         c <- word */
/*     } */
/*     /1* close(c) *1/ */
/* } */
func makeRadoms(c chan int) {
    for {
        c <- rand.Intn(2000)

    }
}
func makeID(idChan chan int) {
    var id int
    id = 0
    for {
        idChan <- id
        id +=1
    }

}
func main () {
    /* wordChannel := make(chan string) */
    /* randoms := make(chan int) */
    /* go makeRadoms(randoms) */
    /* go emit(wordChannel) */
    /* for n := range randoms { */
    /*     fmt.Printf("%d ",n) */

    /* } */

    /* fmt.Printf("\n") */
    /* word := <- wordChannel */
    /* fmt.Printf("%s \n",word) */
    /* word = <- wordChannel */
    /* fmt.Printf("%s \n",word) */
    /* word = <- wordChannel */
    /* fmt.Printf("%s \n",word) */
    /* word = <- wordChannel */
    /* fmt.Printf("%s \n",word) */
    /* /1* var ok bool *1/ */
    /* /1* word,ok = <- wordChannel *1/ */
    /* word,ok := <- wordChannel */
    /* fmt.Printf("%s %t\n",word,ok) */
    idChan := make(chan int)
    go makeID(idChan)
    fmt.Printf("%d \n",<-idChan)
    fmt.Printf("%d \n",<-idChan)
    fmt.Printf("%d \n",<-idChan)
    fmt.Printf("%d \n",<-idChan)
    fmt.Printf("%d \n",<-idChan)
    fmt.Printf("%d \n",<-idChan)
    fmt.Printf("%d \n",<-idChan)
    fmt.Printf("%d \n",<-idChan)
    fmt.Printf("%d \n",<-idChan)
    fmt.Printf("%d \n",<-idChan)
    fmt.Printf("%d \n",<-idChan)
}

