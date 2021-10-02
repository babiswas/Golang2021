package main
import (
   "fmt"
   "strconv"
   "os"
   "bufio"
)


func main(){
   scanner1:=bufio.NewScanner(os.Stdin)
   fmt.Println("Enter the first number")
   scanner1.Scan()
   f1,err:=strconv.ParseFloat(scanner1.Text(),64)
   if err!=nil{
     fmt.Println("Error occured")
   }
   scanner2:=bufio.NewScanner(os.Stdin)
   fmt.Println("Enter the second number")
   scanner2.Scan()
   f2,err:=strconv.ParseFloat(scanner2.Text(),64)
   if err!=nil{
     fmt.Println("Error occured")
   }
   fmt.Printf("The sum of the numbers is %.2f",f1+f2)
}