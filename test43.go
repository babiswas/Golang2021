package main
import (
    "fmt"
    "flag"
)

func main(){
   wordptr:=flag.String("word","foo","a string")
   numptr:=flag.Int("numb",42,"an int")
   boolptr:=flag.Bool("fork",false,"a bool")
   var svar string
   flag.StringVar(&svar,"svar","bar","a string var")
   fmt.Println("word:",*wordptr)
   fmt.Println("numb:",*numptr)
   fmt.Println("fork:",*boolptr)
   fmt.Println("svar:",svar)
   fmt.Println("tail:",flag.Args())
}