package main
import "fmt"

func summer(element ...int){
  sum:=0
  for _,val:=range element{
      sum=sum+val
  }
  fmt.Println("The sum of the numbers is:",sum)
}

func main(){
  arr:=[]int{1,2,3,4,5,6,7}
  fmt.Println("The sum of the array is:")
  summer(arr...)
}