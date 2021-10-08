package main
import (
      "fmt"
      "os"
      "io"
)

func main(){
   if len(os.Args)!=2{
     fmt.Println("Please specify the file name")
     return
   }
   f,err:=os.Open(os.Args[1])
   if err!=nil{
     fmt.Println("Error occured")
     return
   }
   defer f.Close()
   b:=make([]byte,18)
   for n:=0;err==nil;{
       n,err=f.Read(b)
       if err==nil{
         fmt.Print(string(b[:n]))
       }
   }
   if err!=nil && err!=io.EOF{
      fmt.Println("\n\nError:",err)
   }
}