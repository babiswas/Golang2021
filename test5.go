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
   num1,_:=strconv.ParseInt(scanner1.Text(),10,64)
   scanner2:=bufio.NewScanner(os.Stdin)
   fmt.Println("Enter the second number")
   scanner2.Scan()
   num2,_:=strconv.ParseInt(scanner2.Text(),10,64)
   fmt.Printf("The divison of the number is %.2f",float32(num1)/float32(num2)) 
}