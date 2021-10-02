package main
import (
   "fmt"
   "os"
   "bufio"
   "strconv"
)


func main(){
  scanner:=bufio.NewScanner(os.Stdin)
  a:=make([]int64,5)
  fmt.Println(a)
  for i:=0;i<len(a);i++{
    fmt.Println("Enter the value")
    scanner.Scan()
    a[i],_=strconv.ParseInt(scanner.Text(),10,64)
  }
  for _,v:=range a{
     fmt.Println(v)
  }
}