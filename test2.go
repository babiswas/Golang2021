package main
import (
   "fmt"
   "bufio"
   "os"
   "strconv"
)



func main(){
   scanner1:=bufio.NewScanner(os.Stdin)
   scanner2:=bufio.NewScanner(os.Stdin)
   fmt.Println("Enter the first number")
   scanner1.Scan()
   data1,err1:=strconv.ParseInt(scanner1.Text(),10,64)
   if err1!=nil{
     fmt.Println("Error occured")
   }
   fmt.Println("Enter the second number")
   scanner2.Scan()
   data2,err2:=strconv.ParseInt(scanner2.Text(),10,64)
   if err2!=nil{
       fmt.Println("Error occured")
   }
   fmt.Printf("The sum of the number is %d",data1+data2)
}