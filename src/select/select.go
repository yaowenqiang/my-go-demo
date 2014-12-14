package main
import (
    "fmt"
    "time"
)

func emit(wordChannel chan string,done chan bool) {
    defer close(wordChannel)
    words := []string{"The","Quick","Brown","Fox"}
    t := time.NewTimer(3*time.Second)
    i := 0
    for {
        select {
            case wordChannel <- words[i]:
                i +=1
                if i == len(words) {
                    i = 0
                }
            case <-done:
                /* fmt.Printf("Got Done!\n") */
                /* close(done) */
                done <- true
                return
            case <-t.C:
                return

        }

    }

}
func main () {
    wordChan := make(chan string)
    doneCh  := make(chan bool)
    go emit(wordChan,doneCh)
    /* for i := 0;i < 100; i++ { */
    for word := range wordChan{
        fmt.Printf("%s \n",word)

    }
    /* doneCh <- true */
    /* <- doneCh */
}
