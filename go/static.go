package main
import "fmt"
func main(){
    const a1 = 55
    fmt.Println(a1)

    var name = "Sam"
    fmt.Println("type %T value %v",name,name)

    const a = 5
    var intVar int = a
    var int32Var int32 = a
    var float64Var float64 = a
    var complex64Var complex64 = a
    fmt.Println("intVar",intVar, "\nint32Var", int32Var, "\nfloat64Var", float64Var, "\ncomplex64Var",complex64Var)

}
