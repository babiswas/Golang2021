package main
import (
      "fmt"
      "os"
      "path/filepath"
)

func main(){
   if len(os.Args)<2{
     fmt.Println("Please provide path of the directory")
     return
   }
   root,err:=filepath.Abs(os.Args[1])
   if err!=nil{
      fmt.Println("Error :",err)
      return
   }
   fmt.Println("the absolute path of the root is:",root)
   type filedata struct{
     filecount int64
     directorycount int64
   }
   var t filedata=filedata{0,0}
   filepath.Walk(root,func(path string,info os.FileInfo,err error)error{
       if info.IsDir(){
          t.directorycount++
       }else{
          t.filecount++
       }
       fmt.Println(path)
       return nil
   })
   fmt.Printf("Total %d directory and total %d file",t.directorycount,t.filecount)
}