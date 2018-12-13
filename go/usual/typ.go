package main
import "fmt"

const a0 int = 11
const(
    a01 = a0
    a02 = 2
)

var b int = 321

var c = 666

func main(){

    var a int
    a = 65
    a1 := string(a)
    d := 6667
    fmt.Println(a)
    fmt.Println(a0)
    fmt.Println(a01)
    fmt.Println(a02)
    fmt.Println(a1)
    fmt.Println(b)
    fmt.Println(c)
    fmt.Println(d)
    fmt.Println(1 << 10 << 10)

    x := 1
    var p *int = &x
    fmt.Println(p)


}


