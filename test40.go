package main
import (
   "fmt"
   "encoding/json"
)


type Person struct{
  Name string  `json:"name"`
  Id int       `json:"id"`
}

func main(){
  p:=Person{Name:"Bapan",Id:1}
  persons:=[2]Person{Person{Name:"Apan",Id:3},Person{Name:"Bapan",Id:4}}
  data,err:=json.Marshal(p)
  if (err!=nil){
     fmt.Println("error occured")
  }
  fmt.Println(string(data))
  data1,_:=json.Marshal(persons)
  fmt.Println(string(data1))
  data2:=json.RawMessage(string(data))
  bytes,err:=data2.MarshalJSON()
  fmt.Println(bytes)
  if err!=nil{
    fmt.Println("Error occured in marshal json")
  }
  var p1 Person
  err=json.Unmarshal(bytes,&p1)
  if err!=nil{
     fmt.Println("Error Occured")
  }
  fmt.Println(p1)
}