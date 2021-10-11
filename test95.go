package main
import "fmt"
import "bufio"
import "os"
import "strconv"


func main(){
   scanner:=bufio.NewScanner(os.Stdin)
   fmt.Println("Enter value:")
   scanner.Scan()
   val,_:=strconv.ParseInt(scanner.Text(),10,64)
   switch val{
     case 1:
        fmt.Println("First block executed")
        break
     case 2:
        fmt.Println("Second block executed")
        break
     case 3:
        fmt.Println("Third block executed")
        break
     default:
        fmt.Println("DEfault block executed")
        break
   }
}