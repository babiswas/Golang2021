package main
import "fmt"
import "sync"

func main(){
    fmt.Println("Hello World")
    var wg sync.WaitGroup
    wg.Add(1)
    go func(){
      defer wg.Done()
      fmt.Println("Inside go routine")
    }()
    wg.Wait()
    fmt.Println("Finished Execution")
}