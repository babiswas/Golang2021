package main
import "fmt"
import "sync"
import "strconv"

func main(){
   ch1:=make(chan string,10)
   ch2:=make(chan string,10)
   var wg sync.WaitGroup
   fmt.Println("Creating go routines")
   wg.Add(1)
   go func(wg *sync.WaitGroup){
     fmt.Println("Routine 1 executed")
     defer wg.Done()
     for i:=0;i<25;i++{
       ch1<-"hello"+strconv.Itoa(i)
     }
   }(&wg)
   wg.Add(1)
   go func(wg *sync.WaitGroup){
     fmt.Println("Routine 2 executed")
     defer wg.Done()
     for i:=0;i<25;i++{
        ch2<-"bello"+strconv.Itoa(i)
     }
   }(&wg)
   wg.Add(1)
   go func(wg *sync.WaitGroup){
      defer wg.Done()
      for i:=0;i<2;i++{
        select{
        case msg1:=<-ch1:
             fmt.Println(msg1)
             for i:=0;i<24;i++{
                 msg1=<-ch1
                 fmt.Println(msg1)
             }
        case msg2:=<-ch2:
             fmt.Println(msg2)
             for i:=0;i<24;i++{
                 msg2=<-ch2
                 fmt.Println(msg2)
             }
        }
      }
   }(&wg)
   wg.Wait()
   close(ch1)
   close(ch2)
}
