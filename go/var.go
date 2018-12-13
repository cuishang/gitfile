package main
import "fmt"
import "math"
func main(){
    var age int
    fmt.Println("my age is",age)
    age = 18
    fmt.Println("cui age is",age)
    age = 22
    fmt.Println("my new age is",age)
    var name int = 222
    fmt.Println("my name is ",name)
    var add = 333
    fmt.Println("my add is ",add)

    var v1,v2 int = 100,50
    var v3 = v1 * v2
    fmt.Println("the area is ", v3)

    var v4,v5 int
    fmt.Println("v4 is",v4,"v5 is",v5)
    v4 = 400
    v5 = 500
    fmt.Println("new v4 is",v4,"new v5 is",v5)


    var(
	name1 = "ccui"
	age1 = 22
	tall int
    )
    fmt.Println(name1,age1,tall)

    name2,age2 := "cui",24
    fmt.Println("the new is ",name2,age2)

   v7,v8 :=182.3,196.4
   c := math.Min(v7,v8)
   fmt.Println("min is ",c)

}
