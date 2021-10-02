package main
import "fmt"

func main(){
  var a []int=[]int{1,2,3,4,5}
  fmt.Println(a)
  a=append(a,67)
  fmt.Println(a)
  m:=make([]int,4)
  fmt.Println(m)
  for i:=0;i<len(m);i++{
    m[i]=10+i
  }
  fmt.Println(m)
  fmt.Printf("The len of m is %d",len(m))
  fmt.Printf("\n")
  fmt.Printf("the capacity of m is %d",cap(m))
}