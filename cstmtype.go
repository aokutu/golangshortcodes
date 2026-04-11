
package main
import "fmt"

type Age int

func main() {
    var x int = 40
    var a Age

    a = Age(x)

    fmt.Println(a) // 40
}