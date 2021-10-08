package main
import "fmt"

func main(){
   values:=make([]byte,20)
   str1:="hello bello"
   copied:=copy(values,str1)
   fmt.Println(copied)
   fmt.Println(values)
   fmt.Println(string(values))
}