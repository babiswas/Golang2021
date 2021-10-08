package main
import (
   "fmt"
   "os"
)

func main(){
  err:=os.Remove("bello.txt")
  if err!=nil{
     fmt.Println("Error occured")
     return
  }
}