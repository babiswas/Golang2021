package main
import (
    "fmt"
    "os"
    "path/filepath"
)

func main(){
   wd,err:=os.Getwd()
   if err!=nil{
     fmt.Println("Error occured")
   }
   root,err:=filepath.Abs(wd)
   fmt.Printf("The root directory is %s",root)
   fmt.Printf("\n")
   if err!=nil{
     fmt.Println("Error occured")
   }
   fmt.Println("Listing files in ",root)
   type fl struct{
      files int
      dirs int
   }
   var d fl=fl{0,0}
   filepath.Walk(root,func(path string,info os.FileInfo,err error)error{
      if info.IsDir(){
         d.dirs++
      }else{
         d.files++
      }
      return nil
   })
   fmt.Printf("\n")
   fmt.Printf("No of files is %d and no of directories is %d",d.files,d.dirs)
}