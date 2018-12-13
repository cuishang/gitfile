package main
import "fmt"

type Salary interface{
    Ca()int
}

type Permanent struct{
    empid int
    base int
    pf int
}

typr Contract struct{
    empid int
    base int
}

func (p Permanent) Ca()int{
    return p.base + p.pf
}

func (c Contract) Ca()int{
    return c.base
}

func main()



