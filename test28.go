package main
import "fmt"


func modify(x []int){
  x[0]=300
}


func main(){
  var x []int=[]int{3,4,5}
  fmt.Println(x)
  modify(x)
  fmt.Println(x)
}