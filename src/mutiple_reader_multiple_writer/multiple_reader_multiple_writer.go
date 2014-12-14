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
func getter (url string,size chan string) {
    length,err := getPage(url)
    if err == nil {
        size <- fmt.Sprintf("%s has length of %d\n",url,length)
    } else {
        os.Exit(1)
    }
}
func worker (urlCh chan string,sizeCh chan string,id int) {
    for {
        url := <- urlCh
        length,err := getPage(url)
        if err == nil {
            sizeCh <- fmt.Sprintf("%s has length of %d (%d)\n",url,length,id)
        } else {
            sizeCh <- fmt.Sprintf("Error getting %s  %s\n",url,err)
        }

    }
}
func generator (url string,urlCh chan string) {
    urlCh <- url

}
func main () {
    urls := []string{"http://baidu.com/","http://www.sohu.com/","http://www.163.com/","http://www.so.com/","http://www.oschina.net/","http://www.soso.com","http://www.hao123.com"}
    /* for _,url := range urls { */
    /*     pageLength,err := getPage(url) */
    /*     if err != nil { */
    /*         os.Exit(0); */

    /*     } */
    /*     fmt.Printf("%s is length %d \n",url,pageLength) */

    /* } */
    /* size := make(chan string) */

    urlCh := make(chan string)
    sizeCh := make(chan string)

    for i := 0;i<len(urls);i++ {
        go worker(urlCh,sizeCh,i)

    }

    /* for _,url := range urls{ */
    /*     go getter(url,size) */
    /* } */
    /* for i := 0;i<len(urls); i++ { */
    /*     fmt.Printf("%s",<-size) */
    /* } */
    for _,url := range urls {
        /* urlCh <- "http://www.baidu.com" */
        go generator(url,urlCh)
        /* urlCh <- url */
    }
    for i:=0;i<len(urls);i++ {
        fmt.Printf("%s",<-sizeCh);
    }
}


