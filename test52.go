package main
import "fmt"


type Person struct{
  name string
  id int64
}

type Teacher struct{
  name string
  id int64
}

type Human interface{
   GetName() string
}

func (p *Person)GetName()string{
   return p.name
}

func (t *Teacher)GetName()string{
   return t.name
}

func describe_human(h Human){
   switch h.(type){
     case *Person:
         fmt.Println("Person is displayed")
         fmt.Println(h.GetName())
         break
     case *Teacher:
         fmt.Println("Teacher is displayed")
         fmt.Println(h.GetName())
         break
     default:
         fmt.Println("Default block is displayed")
         break
   }
}

func main(){
   p:=Person{"Bapan",12}
   t:=Teacher{"Teacher",13}
   describe_human(&t)
   describe_human(&p)
}