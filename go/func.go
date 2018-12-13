package main
import "fmt"

func reca(length, width float64)(area, perimeter float64) {  
    area = length * width
    perimeter = (length + width) * 2
    return // 不需要明确指定返回值，默认返回 area, perimeter 的值
}

func main(){
    a1,b1 := reca(2.0,4.0)
    fmt.Println("area is %f le is %f",a1,b1)
}
