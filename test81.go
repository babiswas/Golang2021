package main
import "fmt"
import "os/exec"

func main(){
  cmd:=exec.Command("python","--version").Run()
  fmt.Printf("%s\n",cmd)
}