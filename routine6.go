package main
import "runtime"
import "fmt"

func main(){
   numcore:=runtime.NumCPU()
   runtime.GOMAXPROCS(numcore)
   fmt.Println(numcore)
}