package main
import "fmt"
import "os"
import "os/user"
import "strconv"

func main(){
  uid:=os.Getuid()
  u,err:=user.LookupId(strconv.Itoa(uid))
  if err!=nil{
    fmt.Println("Error:",err)
    return
  }
  fmt.Printf("User:%s and userid :%d\n",u.Username,uid)
}