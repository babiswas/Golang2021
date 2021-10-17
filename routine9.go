package main
import "fmt"
import "sync"



func main(){
  var wg sync.WaitGroup
  ch1:=make(chan int64,10)
  ch2:=make(chan int64,10)
  arr:=[]int64{1,2,3,4,5,6,7}
  func1:=func(wg *sync.WaitGroup,data ...int64){
      defer wg.Done()
      for _,v:=range data{
         ch1<-v
      }
      close(ch1)
  }
  brr:=[5]int64{1,2,3,4,5}
  func2:=func(wg *sync.WaitGroup,m [5]int64){
     defer wg.Done()
     for _,v:=range m{
        ch2<-v
     }
     close(ch2)
  }

  func3:=func(wg *sync.WaitGroup){
     var msg int64
     task1:=false
     task2:=false
     defer wg.Done()
     for{
         select{
            case msg=<-ch1:
              fmt.Println("Traversing Channel1")
              if msg!=0{
                fmt.Println(msg)
              }
              for el:=range ch1{
                 fmt.Println(el)
              }
              task1=true
            case msg=<-ch2:
               fmt.Println("Traversing Channel2")
               if msg!=0{
                  fmt.Println(msg)
               }
               for el:=range ch2{
                 fmt.Println(el)
               }
               task2=true
            default:
              fmt.Println("No message")
         }
         if task1 && task2{
            break
         }
     }
  }
  wg.Add(1)
  go func1(&wg,arr ...)
  wg.Add(1)
  go func2(&wg,brr)
  wg.Add(1)
  go func3(&wg)
  wg.Wait()
}
