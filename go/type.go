package main
import "fmt"
func main(){
    a := true
    b := false
    fmt.Println("a:",a,"b:",b)
    c := a && b
    fmt.Println("c:a && b :",c)
    d := a || b
    fmt.Println("d:a || b :",d)

    var e int = 89
    f := 99
    fmt.Println("value e is",e,"value f is ",f)
    var f1 float64 = float64(f)
    fmt.Println("f1 is ",f1)

    g := complex(5,7)
    g1 := 8 + 27i
    add := g + g1
    fmt.Println("sum is ",add)
    
    s1 := "ccui"
    s2 := "ioi"
    s3 := s1 +" " +s2
    fmt.Println("s3 is ",s3)

}
