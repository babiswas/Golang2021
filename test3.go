package main
import (
   "bufio"
   "fmt"
   "os"
   "strconv"
)


func main(){
   scanner1:=bufio.NewScanner(os.Stdin)
   fmt.Println("Enter the principal amount")
   scanner1.Scan()
   principal,err:=strconv.ParseInt(scanner1.Text(),10,64)
   if err!=nil{
     fmt.Println("Error occured")
   }
   scanner2:=bufio.NewScanner(os.Stdin)
   fmt.Println("Enter the rate of intrest")
   scanner2.Scan()
   rate,err1:=strconv.ParseInt(scanner2.Text(),10,64)
   if err1!=nil{
     fmt.Println("Error Occured")
   }
   scanner3:=bufio.NewScanner(os.Stdin)
   fmt.Println("Enter the time")
   scanner3.Scan()
   time,err3:=strconv.ParseInt(scanner3.Text(),10,64)
   if err3!=nil{
      fmt.Println("Error occured")
   }   
   fmt.Printf("The intrest is %.2f",float32((principal*rate*time)/100))
}