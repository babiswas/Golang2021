package main
import "fmt"

func main(){
  var arr []int=[]int{2,3,4,5,4,3,8}
  for i,v:=range arr{
    for j,x:=range arr{
        if i!=j && x==v{
           fmt.Println(x)
        }
    }
  }
}