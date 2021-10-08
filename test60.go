package main
import "fmt"
import "os"
import "bytes"


func main(){
  dst,err:=os.OpenFile("hello2.txt",os.O_CREATE|os.O_WRONLY,0666)
  if err!=nil{
     fmt.Println("Error:",err)
     return
  }
  defer dst.Close()
  str:=[]string{"hello","bello","mello","tello","nello"}
  b := bytes.NewBuffer(make([]byte, 0, 16))
  for _,v:=range str{
     fmt.Fprintf(b,"%s",v)
     b.WriteRune('\n')
     if _,err:=b.WriteTo(dst);err!=nil{
        fmt.Println("Error:",err)
        return
     }
  }
}