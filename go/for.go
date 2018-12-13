package main
import "fmt"

func main(){

    for i := 4;i <= 10;i++{
	fmt.Println("num is %d",i)
    }

    for j := 1;j <= 10;j++ {
	if j > 8{
	    break
	}
	fmt.Printf("nux is %d ",j)
    }
    fmt.Printf("\nline after for loop")
}
