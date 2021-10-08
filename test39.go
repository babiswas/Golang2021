package main
import (
   "fmt"
   "strings"
)

func main(){
  s1:="hello world"
  fmt.Println(s1)
  s1=strings.ToUpper(s1)
  fmt.Println("After to upper the strings value is:")
  fmt.Println(s1)
  fmt.Println("After performing to lower")
  s1=strings.ToLower(s1)
  fmt.Println(s1)
}