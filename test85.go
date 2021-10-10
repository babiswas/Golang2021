package main
import (
  "os/exec"
  "fmt"
)

func main(){
  cmd:=exec.Command("python")
  err:=cmd.Run()
  if err!=nil{
    fmt.Println("Error:",err)
    return
  }
}