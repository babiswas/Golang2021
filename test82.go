package main
import "fmt"
import "os/exec"


func main(){
  cmd:=exec.Command("python","--version")
  if data,err:=cmd.CombinedOutput();err!=nil{
    fmt.Println(err)
    return
  }else{
    fmt.Println(string(data))
  }
  fmt.Println(cmd)
  fmt.Println("cmd:",cmd.Args[0])
  fmt.Println("Args:",cmd.Args[1:])
  fmt.Println("PID:",cmd.Process.Pid)
}