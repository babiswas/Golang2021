package main
import "fmt"

func main(){
  var a[]int=[]int{1,3,4,56,7,12}
  fmt.Println("Displaying in normal for loops")
  for i:=0;i<len(a);i++{
    fmt.Println(a[i])
  }
  fmt.Println("Displaying using range")
  for index,v:=range a{
   fmt.Println(index,v)
  }
  fmt.Println("Using Printf")
  for i,v:=range a{
    fmt.Printf("%d :%d",i,v)
    fmt.Printf("\n")
  }
}