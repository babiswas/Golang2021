package main
import "fmt"

func main(){
  sl1:=make([]int64,10)
  fmt.Println("The length of the slice is :",len(sl1))
  fmt.Println("the capacity of the slice is:",cap(sl1))
  fmt.Println("Populating slice:")
  for i,_:=range sl1{
     fmt.Println("Enter data:")
     fmt.Scanln(&sl1[i])
  }
  fmt.Println("Displaying Slice:")
  for _,v:=range sl1{
     fmt.Println(v)
  }
}