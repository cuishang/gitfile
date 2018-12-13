package main
import "fmt"

func main(){
    a := [...]int{21,23,45,67,87,34,2,1}
    sum := int(0)

    for _,v := range a{
	//fmt.Printf("%d the element of a is %.2d\n",i,v)
	sum +=v
    }
    fmt.Println("\nsum of all element is ",sum)
}
