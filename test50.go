package main
import (
   "fmt"
   "os"
   "io"
)


func CopyFile(from,to string)(int64,error){
   src,err:=os.Open(from)
   if err!=nil{
      return 0,err
   }
   defer src.Close()
   dst,err:=os.OpenFile(to,os.O_WRONLY|os.O_CREATE,0644)
   if err!=nil{
       return 0,err
   }
   defer dst.Close()
   return io.Copy(dst,src)
}


func main(){
   if len(os.Args)<3{
      fmt.Println("Please provide source and destination files")
   }
   CopyFile(os.Args[1],os.Args[2])
}