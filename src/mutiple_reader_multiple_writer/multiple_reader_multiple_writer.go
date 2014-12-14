package main
import (
    "fmt"
    "io/ioutil"
    "net/http"
    "os"
)


func getPage(url string) (int,error){
    resp,err := http.Get(url)
    if err != nil {
        return 0,err
    }
    defer resp.Body.Close()
    body ,err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return 0,err
    }
    return len(body),nil

}
func getter (url string,size chan int) {
    length,err := getPage(url)
    if err == nil {
        size <- length
    } else {
        os.Exit(1)

    }

}
func main () {
    urls := []string{"http://baidu.com/","http://www.sohu.com/","http://www.163.com/"}
    /* for _,url := range urls { */
    /*     pageLength,err := getPage(url) */
    /*     if err != nil { */
    /*         os.Exit(0); */

    /*     } */
    /*     fmt.Printf("%s is length %d \n",url,pageLength) */

    /* } */
    size := make(chan int)
    for _,url := range urls{
        go getter(url,size)
    }
    for i := 0;i<3; i++ {
        fmt.Printf("%d \n",<-size)
    }
}


