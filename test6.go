package main
import "fmt"

func main(){
   var marks=[5]int{99,80,97,43,98}
   sum:=0
   for i:=0;i<5;i++{
     sum=sum+marks[i]
   }
   fmt.Printf("Total marks obtained is %d",sum)
   fmt.Printf("\n")
   fmt.Printf("Average marks is %.2f",float32(sum)/float32(5))
}