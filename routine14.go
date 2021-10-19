package main
import "fmt"
import "sync"

func myFunc(wg *sync.WaitGroup){
   defer wg.Done()
   fmt.Println("Inside my goroutine")
}

func main(){
   var wg sync.WaitGroup
   fmt.Println("Hello World")
   wg.Add(1)
   go myFunc(&wg)
   wg.Wait()
   fmt.Println("Finished execution")
}