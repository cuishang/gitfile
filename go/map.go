package main
import "fmt"

func main(){
    var mapp map[string]int
    if mapp == nil{
	fmt.Println("map is null,going to make one .")
    }else{
	mapp = make(map[string]int)
    }

    mapp1 := make(map[string]int)
    mapp1["china"] = 960
    mapp1["russ"] = 1700
    mapp1["usa"] = 920
    fmt.Println("mapp1 Map is :",mapp1)
    fmt.Println("length of delete is",len(mapp1))
    delete(mapp1,"usa")
    fmt.Println("delect usa mapp1 is ",mapp1)   
    fmt.Println("length of delete is",len(mapp1))
 
    coun := "china"
    fmt.Println("cou is", coun," is big ",mapp1["china"])

    pro := map[string]int{
	"ccui":20,
	"cui":21,
    }
    pro["cccc"] = 24
    for a,b := range pro{
	fmt.Println("pro[%s] = %d\n",a,b)
    }
    newpro := pro
    newpro["cccc"] = 25
    fmt.Println("cccc has been change",pro)

}
