package main
import (
   "fmt"
   "os"
   "time"
   "os/exec"
)


func forkProcess()error{
  cmd:=exec.Command(os.Args[0],"daemon")
  cmd.Stdout,cmd.Stderr,cmd.Dir=os.Stdout,os.Stderr,"/"
  return cmd.Start()
}



func main(){
  fmt.Printf("[%d] start\n",os.Getpid())
  fmt.Printf("[%d] PPID:\n",os.Getpid(),os.Getppid())
  defer fmt.Printf("[%d] Exit \n",os.Getpid())
  if err:=forkProcess();err!=nil{
    fmt.Printf("[%d] Fork error:%s\n",os.Getpid(),err)
    return
  }
}


func runDaemon(){
   for{
     fmt.Printf("[%d] Daemon mode\n",os.Getpid())
     time.Sleep(time.Second *10)
   }
}