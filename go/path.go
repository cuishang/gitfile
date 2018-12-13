package main
import "fmt"

func change(val *int){
    *val = 555
}

func main(){
    b := 255
    var a *int = &b
    fmt.Printf("type of a is %T\n",a)
    fmt.Println("add of a is",a)
    fmt.Println("num of a is",*a)
    *a++
    fmt.Println("num of a is",*a)
    

    c := 25
    var d *int
    if d == nil {
	fmt.Println("d is ",d)
	d = &c
	fmt.Println("d after init is ",d)
    }

    e := 666
    fmt.Println("before change e ",e)
    f := &e
    change(f)
    fmt.Println("after change e ",e)

}
