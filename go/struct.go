package main
import "fmt"

type emm struct{
    fname,lname string
    age,salary int
}

func main(){
    var emp emm
    emp.lname = "emp"
    fmt.Println("emp000 is",emp)

    emp1 := emm{
	fname:"cui",
	lname:"xinaghui",
	age:20,
	salary:2000,
    }
    emp2 := emm{

    }

    emp3 := &emm{"uuu","iiiiii",10,30000}

    fmt.Println("emp2 is ",emp2)
    fmt.Println("emp1 is ",emp1)
    fmt.Println("emp1.name is ",emp1.lname)
    fmt.Println("emp1.age is ",emp1.age)
    fmt.Println("emp3 is ",*emp3)
    fmt.Println("emp3.fname is ",(*emp3).fname)
    fmt.Println("emp3.lname is ",emp3.lname)

    per := struct {
	name string
	age int
    }{
	name:"ccccccccccccc",
	age:20,
    }
    fmt.Println("per is ",per) 
    

}
