package main
import "fmt"


type Point struct{
  x int64
  y int64
}

type Circle struct{
  radius float64
  *Point
}


func main(){
  c1:=Circle{4.67,&Point{3,4}}
  fmt.Println(c1)
  fmt.Println(c1.x)
  fmt.Println(c1.y)
}


