package main
import "fmt"
import "bufio"
import "os"

type Service interface{
   service() string
}

type DatabaseService struct{
  database string
}

type UserService struct{
   username string
}

func (d *DatabaseService)service()string{
   return d.database
}

func (u *UserService)service()string{
   return u.username
}

func main(){
   var key string
   servicemap:=make(map[string]Service)
   servicemap["sql"]=&DatabaseService{"SQL Service"}
   servicemap["hello"]=&UserService{"user1"}
   fmt.Println("The length of the map is:",len(servicemap))
   for{
      scanner:=bufio.NewScanner(os.Stdin)
      fmt.Println("Enter the key")
      scanner.Scan()
      key=scanner.Text()
      fmt.Println("Displaying Service")
      fmt.Println(servicemap[key].service())
   }
}

