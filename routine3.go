package main
import "fmt"
import "time"

func main(){
  t:=time.Now()
  time.Sleep(time.Second*2)
  arr:=[10]int64{1,2,3,4,5,6,7,8,9,10}
  var sum int64
  sum=0
  for i:=0;i<len(arr);i++{
    sum=sum+arr[i]
  }
  fmt.Println("Time elapsed is:",time.Since(t))
  fmt.Println("The sum of the numbers is:",sum)
}