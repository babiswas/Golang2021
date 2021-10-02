package main
import "fmt"

type Teacher struct{
   name string
   subject string
}

type Professor struct{
   name string
   subject string
   papers int
}

func (p *Professor)teach(){
   fmt.Println("Teaching "+p.subject)
}

func (t *Teacher)teach(){
  fmt.Println("Teaching "+t.subject)
}


type instructor interface{
  teach()
}


func main(){
  t:=Teacher{"X","Physics"}
  p:=Professor{"Y","Electromagnetics",3}
  faculty:=[]instructor{&t,&p}
  for _,f:=range faculty{
    f.teach()
  }
}