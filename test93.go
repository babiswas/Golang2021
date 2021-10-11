package main
import "fmt"
import "time"


func f(n int){
  fmt.Println("Executing routine:",n)
  for i:=0;i<10;i++{
    fmt.Println(n,":",i)
    time.Sleep(time.Second*3)
  }
}

func main(){
  for i:=0;i<10;i++{
    go f(i)
  }
  var input string
  fmt.Scanln(&input)
}