package main
import "fmt"


func test1(x string) func(){
   return func(){
     fmt.Println(x)
   }
}

func main(){
   x:=test1("hello")
   x()
}