package main
import "encoding/json"
import "fmt"

type Person struct{
   Name string `json:"name"`
   Id int64 `json:"id"`
}

func main(){
   p:=Person{"Bapan",12}
   data,_:=json.Marshal(p)
   fmt.Println(string(data))
   data1:=json.RawMessage(data)
   bte,err:=data1.MarshalJSON()
   if err!=nil{
      fmt.Println("Error occured")
      return
   }
   var p2 Person
   err=json.Unmarshal(bte,&p2)
   fmt.Println(p2)
}