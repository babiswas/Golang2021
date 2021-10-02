package main
import (
    "fmt"
    "os"
    "bufio"
    "strconv"
)

func main(){
   scanner:=bufio.NewScanner(os.Stdin)
   fmt.Println("Enter the number")
   scanner.Scan()
   data,err:=strconv.Atoi(scanner.Text())
   if err!=nil{
     fmt.Println("Error occured")
   }
   fmt.Println(data)
}