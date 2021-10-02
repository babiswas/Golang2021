package main
import (
   "fmt"
   "bufio"
   "strconv"
   "os"
)


type Book struct{
  bookname string
  authorname string
  bookid int64
}


func main(){
  var books [5]Book
  scanner1:=bufio.NewScanner(os.Stdin)
  scanner2:=bufio.NewScanner(os.Stdin)
  scanner3:=bufio.NewScanner(os.Stdin)
  for i:=0;i<len(books);i++{
    fmt.Println("Enter bookname")
    scanner1.Scan()
    books[i].bookname=scanner1.Text()
    fmt.Println("Enter authorname")
    scanner2.Scan()
    books[i].authorname=scanner2.Text()
    fmt.Println("Enter bookid")
    scanner3.Scan()
    books[i].bookid,_=strconv.ParseInt(scanner3.Text(),10,64)
  } 

  for i:=0;i<len(books);i++{
    fmt.Println(books[i])
  }
}