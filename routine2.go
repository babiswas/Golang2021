package main
import "fmt"
import "strconv"

func main(){
   num:=100
   str1:="hello "+strconv.Itoa(num)
   fmt.Println(str1)
   num2:="200"
   num3,_:=strconv.Atoi(num2)
   fmt.Println(num3)
}