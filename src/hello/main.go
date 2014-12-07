package main

import (
	"fmt"
)
const (
    /* message="The answer to life is %d\n" */
    /* answer=42 */

    message="%d %d \n"

    /* answer1=iota */
    answer1=iota*2
    answer2
)

func main() {
    /*var message string*/
    /*message = "hello world,i am message"*/
    /* message:="Hello world i am message 2" */
    /* message:="The answer to life is %d\n" */
    /* answer := 42 */
    /* answer += 1 */
    /* var pi float64 = 3.14 */
    /* var pi float32 = 3.14 */
    pi := float64(3.14)
	fmt.Printf(message,answer1,answer2)
	fmt.Printf("%f\n",pi)
	fmt.Printf("%.2f\n",pi)
    nine := 9
    six := uint(6)
    /* six := uint8(6) */
    /* six := uint32(6) */
    /* six := uint64(6) */
    fmt.Printf("Value is %d\n",nine)
    fmt.Printf("Value is %d\n",six)
    /* isTrue := true */
    /* var isTrue bool */
    isTrue := !true
    b := byte(65)
    fmt.Printf("Value is %t\n",isTrue)
    fmt.Printf("Value is %b\n",b)
    fmt.Printf("Value is %x\n",b)
}
