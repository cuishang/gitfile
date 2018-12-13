package main
import (
    "fmt"
    "os"
)
func main(){
    f,err := os.Open("map.go")
    if err != nil{
	fmt.Println(err)
	return
	}
    fmt.Println(f.Name(),"open success")
}
