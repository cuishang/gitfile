package main
import "fmt"

func call(price int,no int) int{
    var toll = price * no
    return toll
}

func area(l,h float64)(float64,float64){
    var ar = l * h
    var le = (l + h)*2
    return ar,le
}

func area1(s1,s2 float64)(ar1,le1 float64){
    var ar1 = s1 * s2
    var le1 = (s1 + s2)*2
    return
}

func main(){

    p, n := 2, 8
    cal1 := call(p,n)
    fmt.Println("tol is ",cal1)
    
    a,b := area(20.0,4.0)
    fmt.Println("area is %f le is %f",a,b)

    a1,b1 := area1(2.0,4.0)
    fmt.Println("area is %f le is %f",a1,b1)
}
