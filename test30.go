package main
import "fmt"


type Point struct{
   x int64
   y int64
}

func changeX(p1 *Point){
  p1.x=12
  p1.y=13
}


func main(){
   p1:=&Point{x:10,y:11}
   fmt.Println(p1)
   changeX(p1)
   fmt.Println(*p1)
}