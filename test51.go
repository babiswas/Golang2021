package main
import (
   "fmt"
   "os"
   "io"
)


func CopyFile(from,to string)(int64,error){
   fd1,err:=os.Open(from)
   defer fd1.Close()
   if err!=nil{
     fmt.Println("Error occured")
     return 0,err
   }
   fd2,err:=os.OpenFile(to,os.O_WRONLY|os.O_CREATE,0644)
   if err!=nil{
     fmt.Println("Error occured")
     return 0,err
   }
   defer fd2.Close()
   return io.Copy(fd2,fd1)
}


func main(){
   if len(os.Args)<3{
      fmt.Println("Please provide input files")
      return
   }
   CopyFile(os.Args[1],os.Args[2])
   fmt.Println("Operation Complete")
}