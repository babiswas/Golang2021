package main
import "fmt"

func f(data int){
  fmt.Println("Executing:",data)
  for i:=0;i<10;i++{
    fmt.Println(data,";",i)
  }
}

func main(){
  for i:=0;i<10;i++{
    go f(i)
  }
  var input string
  fmt.Scanln(&input)
}