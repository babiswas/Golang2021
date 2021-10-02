package main
import "fmt"

type Shape interface{
  area() int
}

type Rect struct{
   x int
   y int
}

type Square struct{
   x int
}


func (s *Square)area() int{
   return s.x*s.x
}

func (s *Rect) area() int{
   return s.x*s.y
}


func main(){
   s1:=Square{4}
   r1:=Rect{3,4}
   myarea:=[]Shape{&s1,&r1}
   fmt.Println("displaying area")
   for _,v:=range myarea{
     fmt.Println(v.area())
   }
}

