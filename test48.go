package main
import (
   "fmt"
   "bytes"
)

func main(){
  b:=bytes.NewBuffer(make([]byte,26))
  var text=[]string{"hello people","how are you","hello brother"}
  for i:=range text{
     b.Reset()
     b.WriteString(text[i])
     fmt.Printf("Length: %d Capacity: %d",b.Len(),b.Cap())
  }
}