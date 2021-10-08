package main
import (
      "fmt"
      "os"
      "bufio"
      "io"
)

func main(){
  if len(os.Args)!=2{
    fmt.Println("Please specify the path")
    return
  }
  f,err:=os.Open(os.Args[1])
  if err!=nil{
    fmt.Println("Error:",err)
    return
  }
  defer f.Close()
  r:=bufio.NewReader(f)
  var count int
  for err==nil{
    var b []byte
    for moar:=true;err==nil && moar;{
       b,moar,err=r.ReadLine()
       if err==nil{
         fmt.Print(string(b))
       }
       if err == nil{
          fmt.Println()
          count=count+1
       }
    }
  }
  if err!=nil && err!=io.EOF{
    fmt.Println("Error:",err)
    return
  }
  fmt.Println("Row count:",count)
}