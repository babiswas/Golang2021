package main
import "fmt"
import "os/exec"
import "os"

func main(){
  cmd:=exec.Command("python")
  cmd.Stdout=os.Stdout
  cmd.Stderr=os.Stderr
  err:=cmd.Run()
  if err!=nil{
     fmt.Println("Error:",err)
     return
  }
}