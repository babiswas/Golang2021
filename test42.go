package main
import (
   "fmt"
   "os"
)

func main(){
  var1:=os.Args
  var2:=os.Args[1:]
  fmt.Println(var1)
  fmt.Println(var2)
  for i:=0;i<len(var1);i++{
    fmt.Println(var1[i])
  }
}