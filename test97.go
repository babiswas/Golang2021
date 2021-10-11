package main
import "fmt"
import "math"

type Circle struct{
  r int64
}

func (c *Circle)area(radius int64)float64{
  c.r=radius
  return float64(math.Pi*float64(c.r)*float64(c.r))
}

func main(){
  c1:=Circle{4}
  fmt.Println(c1.area(5))
  fmt.Println(c1.area(6))
}