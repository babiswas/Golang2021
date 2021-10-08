package main
import "fmt"
import "os"
import "io"

func main(){
  if len(os.Args)<2{
    fmt.Println("Please provide the path")
    return
  }
  f,err:=os.Open(os.Args[1])
  if err!=nil{
    fmt.Println("Error:",err)
    return
  }
  defer f.Close()
  b:=make([]byte,16)
  for n:=0;err==nil;{
    n,err=f.Read(b)
    if err==nil{
       fmt.Println(string(b[:n]))
    }
  }
  if err!=nil && err!=io.EOF{
    fmt.Println("Error:",err)
  }
}
