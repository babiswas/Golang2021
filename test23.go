package main
import "fmt"

func main(){
   var m map[string]int=map[string]int{"bapan":1,"priya":2,"juli":4}
   fmt.Println(m)
   for key,value:=range m{
     fmt.Println(key,value)
   }
   delete(m,"bapan")
   fmt.Println(m)
}