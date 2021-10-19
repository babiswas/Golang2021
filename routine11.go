package main
import "fmt"
import "net/http"
import "io/ioutil"

func main(){
  client:=http.Client{}
  req,err:=http.NewRequest("GET","https://captivateprimestage1.adobe.com/primeapi/v2/badges",nil)
  req.Header=http.Header{"Accept":[]string{"application/vnd.api+json"},"Authorization":[]string{"oauth c6f8fa81c21621c533cab07af725221e"}}
  res,err:=client.Do(req)
  if err!=nil{
     fmt.Println("Error occured",err)
     return
  }else{
     n,_:=ioutil.ReadAll(res.Body)
     data1:=string(n)
     fmt.Println(data1["links"])
  }
}