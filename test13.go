package main
import "fmt"


type Student struct{
 name string
 studentid int
}

func (s Student) getName() string{
  return s.name
}

func (s Student) getId() int{
   return s.studentid
}

func (s *Student) setId(id int){
  s.studentid=id
}

func (s *Student) setName(name string){
   s.name=name
}


func main(){
   var s1 Student=Student{name:"Bapan",studentid:32}
   var s2 Student=Student{name:"Tapan",studentid:33}
   fmt.Println(s1.getName())
   fmt.Println(s1.getId())
   fmt.Println(s2.getName())
   fmt.Println(s2.getId())
   fmt.Println(s1)
   s1.setId(32)
   s1.setName("Madan")
   fmt.Println(s1)
}