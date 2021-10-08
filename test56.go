package main
import (
  "fmt"
  "os"
  "io/ioutil"
)

func main(){
  if len(os.Args)!=3{
    fmt.Println("Please specify path")
    return
  }
  if err:=ioutil.WriteFile(os.Args[1],[]byte(os.Args[2]),0644);err!=nil{
     fmt.Println("Error:",err)
  }
}