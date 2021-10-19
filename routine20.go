package main
import "fmt"
import "math/rand"
import "time"


func CalculateValue(c chan int){
  value:=rand.Intn(10)
  fmt.Println("Calculated random value:{}",value)
  time.Sleep(1* time.Millisecond)
  c<-value
  fmt.Println("Execute after another go routine perform recieve")
}

func main(){
   fmt.Println("Go channel tutorial")
   valueChannel:=make(chan int)
   defer close(valueChannel)
   go CalculateValue(valueChannel)
   go CalculateValue(valueChannel)
   values:=<-valueChannel
   fmt.Println(values)
}