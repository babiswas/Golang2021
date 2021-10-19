package main
import "fmt"


func main(){
  p:=make(map[string]int)
  fmt.Println(len(p))
  p["hello"]=1
  p["bello"]=2
  p["mello"]=3
  fmt.Println(p)
  for key,value:=range p{
     fmt.Printf("%s:%d",key,value)
     fmt.Println("\n")
  }
  delete(p,"mello")
  fmt.Println(p)
}