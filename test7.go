package main
import "fmt"

func main(){
  test:=[5]int{1,2,3,4,5}
  for i:=0;i<len(test);i++{
    fmt.Println(test[i])
  }
  fmt.Printf("The length of test array is %d",len(test))
}