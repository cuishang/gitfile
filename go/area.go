package main
import "fmt"

func local(num[5]int){
    num[0] = 55
    fmt.Println("inside func ",num)
}

func main(){

    num := [...]int{5,6,7,8,8}
    fmt.Println("函数未处理前:",num)
    fmt.Println("length og num ",len(num))
    local(num)
    fmt.Println("函数处理后:",num)

    var a[3]int
    fmt.Println(a)

    var b[3]int  //b[3]int{12,23,345,4}
    b[0] = 12
    b[1] = 24
    b[2] = 36
    fmt.Println(b)

    c := [...]int{12,24,25,34,34,35,465,567,5,75,76,57}
    for i:=0;i<len(c);i++{
	fmt.Printf("%dth of c is %.2d\n",i,c[i])
	}
    fmt.Println(c)

    d := [4]int{12}
    fmt.Println(d)
 
    e := [...]int{23,34,56,78,342,132}
    fmt.Println(e)

    f := [...]string{"usa","china","23"}
    f1 := f
    f1[0] = "japan"
    fmt.Println("f is ",f)
    fmt.Println("f1 is ",f1)
}
