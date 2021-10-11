package main
import "fmt"
import "time"


func main(){
   c1:=make(chan string)
   c2:=make(chan string)

   go func(){
     for{
      time.Sleep(time.Second*1)
      c1<-"Audio data"
     }
   }()

   go func(){
     for{
       time.Sleep(time.Second*1)
       c2<-"Video data"
     }
   }()

   for{
      select{
        case msg1:=<-c1:
           fmt.Println(msg1)
           break
        case msg2:=<-c2:
           fmt.Println(msg2)
           break
      }
   }

   var input string
   fmt.Scanln(&input)
}