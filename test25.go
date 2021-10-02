package main
import "fmt"


func add(x,y int)(int,int){
   defer fmt.Println("Operation Completed")
   fmt.Println("Performing the operation")
   return x+y,x-y
}

func main(){
   a,b:=add(2,3)
   c,d:=add(5,6)
   fmt.Println(a,b)
   fmt.Println(c,d)
}