package main
import "fmt"

type Student struct{
  name string
  studentid int
}

func main(){
  var s1 Student=Student{name:"Bapan",studentid:32}
  var s2 Student=Student{name:"Tapan",studentid:44}
  fmt.Println(s1)
  fmt.Println(s2)
}