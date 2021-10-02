package main
import "fmt"

func main(){
 var x [5]int=[5]int{1,2,3,4,5}
 var y []int=x[1:3]
 var z []int=x[:4]
 fmt.Println(x)
 fmt.Println(y)
 fmt.Println(z)
 fmt.Printf("The length of y is %d",len(y))
 fmt.Println("\n")
 fmt.Printf("The capacity of y is %d",cap(y))
 fmt.Printf("\n")
 fmt.Println(y[:cap(y)])
}