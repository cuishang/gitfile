package main
import "fmt"
import "math"

type emm struct{
    name string
    age int
    add string
}

type Re struct{
    len int
    wid int
}

type Ci struct{
    r float64
}


func (e emm)printx(){
    fmt.Printf("name is %s age is %d add is %s\n",e.name,e.age,e.add)
}

func (r Re) area()int{
    return r.len*r.wid
}

func (c Ci) area()float64{
    return math.Pi*c.r*c.r
}

func main(){
    emp1 := emm{
	name: "ccui",
 	age:20,
	add:"henan",
    }
    emp1.printx()

    c := Ci{
	r:12,
    }
    fmt.Println("Area of Ci is ",c.area())




}
