package main
import "fmt"
import "os"

func main(){
   if len(os.Args)<2{
     fmt.Println("Please provide path")
     return
   }
   fmt.Println(os.Stat(os.Args[1]))
}