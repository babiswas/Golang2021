package main
import "fmt"
import "os"
import "bufio"


func main(){
  if len(os.Args)<2{
    fmt.Println("Please provide path")
    return
  }
  scanner:=bufio.NewScanner(os.Stdin)
  fmt.Println("Enter data")
  scanner.Scan()
  data:=scanner.Text()
  err:=os.WriteFile(os.Args[1],[]byte(data),0664)
  if err!=nil{
    fmt.Println("Error occured:",err)
    return
  }  
}