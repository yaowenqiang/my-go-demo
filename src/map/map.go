package main
import(
    "fmt"
)

func main () {
    dayMounths := make(map[string]int)
    dayMounths["Jan"] = 31
    dayMounths["Feb"] = 28
    dayMounths["Mar"] = 31
    dayMounths["Apr"] = 30

    dayMounths["May"] = 31
    dayMounths["Jun"] = 30
    dayMounths["Jul"] = 31
    dayMounths["Agu"] = 31

    dayMounths["Sep"] = 30
    dayMounths["Oct"] = 31
    dayMounths["Nov"] = 30
    dayMounths["Dec"] = 31

    fmt.Printf("Days in February: %d \n",dayMounths["Feb"])

}
