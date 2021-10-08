package main
import "fmt"
import "io/ioutil"
import "os"


func main(){
   if len(os.Args)<2{
     fmt.Println("Please provide the path")
     return
   }
   b,err:=ioutil.ReadFile(os.Args[1])
   if err!=nil{
     fmt.Println("Error occured:",err)
   }
   fmt.Println(string(b))
}