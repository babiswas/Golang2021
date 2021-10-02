package main
import (
   "fmt"
   "bufio"
   "strconv"
   "os"
)


type Book struct{
  bookname string
  bookid int64
}

func main(){
   scanner1:=bufio.NewScanner(os.Stdin)
   scanner2:=bufio.NewScanner(os.Stdin)
   var book [3]Book
   for i:=0;i<3;i++{
     fmt.Println("Enter bookname")
     scanner1.Scan()
     book[i].bookname=scanner1.Text()
     fmt.Println("Enter book id")
     scanner2.Scan()
     book[i].bookid,_=strconv.ParseInt(scanner2.Text(),10,64)
   } 
   fmt.Println(book) 
}