package main
import (
   "fmt"
   "bufio"
   "strconv"
   "os"
)


func add(myfunc func()){
   myfunc()
}


func main(){
   test:=func(){
     defer fmt.Println("Hello World")
     scanner:=bufio.NewScanner(os.Stdin)
     fmt.Println("Enter the number")
     scanner.Scan()
     num,_:=strconv.ParseInt(scanner.Text(),10,64)
     fmt.Println("Enter the second number")
     scanner.Scan()
     num1,_:=strconv.ParseInt(scanner.Text(),10,64)
     fmt.Printf("The sum of the number is %d",num+num1)
     fmt.Printf("\n")
   }
   test()
   fmt.Println("Running 2nd implementation")
   add(test)
}