package main
import(
	"fmt"
//	"unicode/utf8"
)

func printx(s string){
    for i:=0;i<len(s);i++ {
	fmt.Printf("%x ",s[i])
    }
}

func printx1(s string){
    runes := []rune(s)
    for i:=0;i<len(runes);i++ {
	fmt.Printf("%c ",runes[i])
    }
    //fmt.Printf("length of %s is %d\n",runes,utf8.RuneCountInString(runes))
}

func putrange(s string){
    for index,rune := range s {
	fmt.Printf("%c start at byte %d\n",rune,index)
    }
}

func main(){
    name := "hello world !"
    fmt.Println(name)
    printx(name)
    println("\n")
    printx1(name)

    putrange(name)
}
