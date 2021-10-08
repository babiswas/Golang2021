package main
import (
     "fmt"
     "os"
)

func main(){
   wd,err:=os.Getwd()
   if err!=nil{
      fmt.Println("Error Occured")
      fmt.Println(err)
   }
   fmt.Printf("Working directory is %s",wd)
   if err=os.Chdir("C:\\Users\\babiswas\\Desktop\\mux");err!=nil{
      fmt.Println("error occured")
      fmt.Println(err)
   }
   if wd,err=os.Getwd();err!=nil{
      fmt.Println(err)
   }
   fmt.Println("\nFinal dir:",wd)
}

/*
os.Getwd()
os.Chdir()
*/