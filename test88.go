package main
import "fmt"
import "os"
import "os/exec"

func main(){
  cmd:=exec.Command("go","run","test11.go")
  cmd.Stdout=os.Stdout
  cmd.Stderr=os.Stderr
  cmd.Stdin=os.Stdin
  err:=cmd.Run()
  if err!=nil{
     fmt.Println("Error:",err)
     return
  }
}