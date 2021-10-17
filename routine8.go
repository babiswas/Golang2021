  package main
import "fmt"
import "sync"

var count=0

func counter(wg *sync.WaitGroup,m *sync.Mutex){
   defer wg.Done()
   m.Lock()
   defer m.Unlock()
   count=count+1
}
func main(){
  var wg sync.WaitGroup
  var m sync.Mutex
  for i:=0;i<1000;i++{
     wg.Add(1)
     go counter(&wg,&m)
  }
  wg.Wait()
  fmt.Println(count)
}


