package main
import (
    "fmt"
)

func emit(wordChannel chan string,done chan bool) {
    words := []string{"The","Quick","Brown","Fox"}
    i := 0
    for {
        select {
            case wordChannel <- words[i]:
                i +=1
                if i == len(words) {
                    i = 0
                }
            case <-done:
                fmt.Printf("Got Down!\n")
                close(done)
                return

        }

    }

}
func main () {
    wordChan := make(chan string)
    doneCh  := make(chan bool)
    go emit(wordChan,doneCh)
    for i := 0;i < 100; i++ {
        fmt.Printf("%s \n",<-wordChan)

    }
    doneCh <- true
}
