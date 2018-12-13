package main
import "fmt"

func main(){

}

func Add(a ...int)(a,b,c int){
    a,b,c = 1,2,3
    fmt.Println(a)
}
