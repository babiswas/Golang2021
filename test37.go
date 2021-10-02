package main
import (
   "fmt"
   "bufio"
   "os"
   "strconv"
)

type Person struct{
   name string
   id int64
}

type Student struct{
   person Person
   marks [5]int64
}


func (s *Student) addNameId(){
    scanner:=bufio.NewScanner(os.Stdin)
    fmt.Println("Enter the name of the student")
    scanner.Scan()
    s.person.name=scanner.Text()
    fmt.Println("Enter the id of the student")
    scanner.Scan()
    s.person.id,_=strconv.ParseInt(scanner.Text(),10,64)
}

func (s *Student)addMarks(){
   scanner:=bufio.NewScanner(os.Stdin)
   for i:=0;i<5;i++{
      fmt.Println("Enter the marks")
      scanner.Scan()
      s.marks[i],_=strconv.ParseInt(scanner.Text(),10,64)
   }
}

func (s Student) getAverage() float64{
   var sum int64
   sum=0
   for _,v:=range s.marks{
      sum=sum+v
   }
   return float64(float64(sum)/float64(len(s.marks)))
}

func main(){
   s1:=Student{}
   s1.addMarks()
   s1.addNameId()
   fmt.Println(s1)
   fmt.Printf("The average marks is %.2f",s1.getAverage())
}