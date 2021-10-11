package main
import "fmt"
import "flag"


func main(){
  max1:=flag.Int("max",6,"the max value")
  flag.Parse()
  fmt.Println(*max1)
}