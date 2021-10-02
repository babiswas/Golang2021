package main
import "fmt"

func main(){
  var x []int=[]int{1,2,3,4,5}
  y:=x
  fmt.Println(x)
  fmt.Println(y)
  y[0]=100
  fmt.Println(x)
  fmt.Println(y)
  m:=map[string]string{"bello":"tello","mello":"kello"}
  fmt.Println(m)
  for key,value:=range m{
    fmt.Printf("%d :%d",key,value)
    fmt.Println("\n")
  }
}