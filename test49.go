package main
import "fmt"

func main(){
  b:=[5]int{1,2,3,4,5}
  for i,v:=range b{
    fmt.Println(i)
    fmt.Println(v)
  }
}