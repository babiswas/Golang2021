package main
import "bytes"
import "fmt"
import "os"
import "os/exec"


func main(){
  b:=bytes.NewBuffer(nil)
  cmd:=exec.Command("python")
  cmd.Stdin=b
  cmd.Stdout=os.Stdout
  fmt.Fprintf(b,"--version",b)
  if err:=cmd.Start();err!=nil{
     fmt.Println(err)
     return
  }
  cmd.Wait()
}