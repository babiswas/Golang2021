package main
import "fmt"
import "log"
import "net/http"
import "sync"

var urls=[]string{"https://google.com","https://twitter.com","https://tutorialedge.net"}

func fetch(url string,wg *sync.WaitGroup)(string,error){
  defer wg.Done()
  resp,err:=http.Get(url)
  if err!=nil{
    fmt.Println(err)
    return "",err
  }
  fmt.Println(resp.Status)
  return resp.Status,nil
}

func homePage(w http.ResponseWriter,r *http.Request){
   fmt.Println("Homepage end point hit")
   var wg sync.WaitGroup
   for _,url:=range urls{
      wg.Add(1)
      go fetch(url,&wg)
   }
   wg.Wait()
   fmt.Println("Returning Response")
   fmt.Fprintf(w,"Responses")
}

func handleRequest(){
   http.HandleFunc("/",homePage)
   log.Fatal(http.ListenAndServe(":8081",nil))
}

func main(){
   handleRequest()
}