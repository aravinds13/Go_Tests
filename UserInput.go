//get user input

package main
import "fmt"

func main(){
  val:=getVal()
  fmt.Println("Double of ",val,"is ",val*2)
}

func getVal() float64{ //func function_name return_type
  var value float64
  fmt.Print("Enter a value: ")
  fmt.Scanf("%f",&value)
  return value
}
