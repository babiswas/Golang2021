package main
import "fmt"
import "encoding/json"

type Teacher struct{
  Name string `json:"name"`
  Id int64   `json:"id"`
  Subject string `json:"subject"`
}

func main(){
  t:=Teacher{"Bapan",23,"Physics"}
  fmt.Println(t)
  data,_:=json.Marshal(t)
  fmt.Println(string(data))
  var t1 Teacher
  data1:=json.RawMessage(string(data))
  byte12,err:=data1.MarshalJSON()
  if err!=nil{
    fmt.Println("Error:",err)
    return
  }
  err=json.Unmarshal(byte12,&t1)
  if err!=nil{
    fmt.Println("Error:",err)
    return
  }
  fmt.Println(t1)
}