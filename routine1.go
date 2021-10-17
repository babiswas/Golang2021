package main
import "fmt"
import "sync"


func main(){
  tasklist:=[5]string{"task1","task2","task3","task4","task5"}
  var wg sync.WaitGroup
  for i:=0;i<5;i++{
    wg.Add(1)
    go f1(tasklist[i],&wg)
  }
  wg.Wait()
}

func f1(str1 string,wg *sync.WaitGroup){
  defer wg.Done()
  for i:=0;i<500;i++{
    fmt.Println(str1)
  }
}