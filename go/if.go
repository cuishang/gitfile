package main
import "fmt"
func main(){
    num := 101
    if num % 2 == 0{
	fmt.Println("the num is even")
    }else{
	fmt.Println("the num is odd")
    }

    if nu := 10; nu % 2 == 0 { //nu范围仅限if代码块
        fmt.Println(nu,"is even") 
    }  else {
        fmt.Println(nu,"is odd")
    }

    no := 999
    if no <= 50 {
	fmt.Println("the less 50")
    }else if no >= 51 && no <= 100{
	fmt.Println("the between 51 and 100")
    }else{
	fmt.Println("the lager 100")
    }


}
