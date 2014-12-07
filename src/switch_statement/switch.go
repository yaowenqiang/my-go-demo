package main
import ( 
    "fmt"
    "os"

)

func main (){
    n,err := fmt.Printf("hello world!\n");
    n = 0
    switch{
        case err != nil:
            os.Exit(1)
        case n == 0:
            fmt.Printf("No bytes output\n")
            fallthrough
        case n != 13:
            fmt.Printf("Wrong number of characters! %d",n)
        default:
            fmt.Printf("OK!")
    }
    fmt.Printf("\n")
    atoz := "the quick brown fox jumps over the lazy dog"
    vowels := 0
    consonants := 0
    for _,f := range atoz {
        switch f {
            case 'a','e','i','o','u':
                vowels +=1
            default:
                consonants +=1
        }
    }
    fmt.Printf("Vowels: %d,Consonants: %d\n",vowels,consonants)
}

