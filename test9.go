package main
import (
   "fmt"
   "os"
   "bufio"
   "strconv"
)

func main(){
   var arr=[5]int64{}
   scanner:=bufio.NewScanner(os.Stdin)
   for i:=0;i<len(arr);i++{
     fmt.Println("Enter the  number")
     scanner.Scan()
     arr[i],_=strconv.ParseInt(scanner.Text(),10,64)
   }
   fmt.Println(arr)
}