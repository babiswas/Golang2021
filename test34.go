package main
import (
   "fmt"
   "bufio"
   "strconv"
   "os"
)

func main(){
  scanner:=bufio.NewScanner(os.Stdin)
  fmt.Println("Enter the value")
  scanner.Scan()
  ans,_:=strconv.ParseInt(scanner.Text(),10,64)
  switch ans{
     case 1:
       fmt.Println("First case executed")
       break
     case 2:
       fmt.Println("Second case executed")
       break
     case 3:
       fmt.Println("Third case executed")
       break
     case 10:
       fmt.Println("Case 10 executed")
       break
     default:
       fmt.Println("default case executed")
  }
}