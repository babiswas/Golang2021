package main
import (
    "fmt"
    "os"
    "bufio"
)

func main(){
   fmt.Println("Enter the string:")
   scanner:=bufio.NewScanner(os.Stdin)
   scanner.Scan()
   data:=scanner.Text()
   fmt.Println("The bytes array for this string is")
   data1:=[]byte(data)
   fmt.Println(data1)
   fmt.Println("The string equivalent of this bytes:")
   fmt.Println(string(data1))
}