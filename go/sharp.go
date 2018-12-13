package main
import "fmt"

func main(){
    a := [...]int{23,34,45,99,65,45,56,34,3}
    var b[]int = a[1:5]
    fmt.Println("b is",b)

    c := []int{6,7,8}
    fmt.Println(c)

    darr := [...]int{54,342,234,34,45,65,6}
    sharp := darr[2:5]
    fmt.Println("修改前：",darr)
    for i := range sharp{
	sharp[i]++
    }
    fmt.Println("修改后：",darr)
}
