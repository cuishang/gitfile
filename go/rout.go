package main
import "fmt"
import "time"

func hello(){
    fmt.Println("hello is routine")
}

func num(){
    for i:=1;i<=5;i++{
	time.Sleep(250 * time.Millisecond)
	fmt.Printf("%d ",i)
    }
}

func alp(){
    for i:='a';i<='e';i++{
	time.Sleep(400 * time.Millisecond)
	fmt.Printf("%c ",i)
    }
}

func main(){
//    go hello()
//    time.Sleep(1 * time.Second)

    go num()
    go alp()
    time.Sleep(3000 * time.Millisecond)
    fmt.Println("\n")
    fmt.Println("main function")
}
