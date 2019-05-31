//functions and returning multiple values from them

package main
import "fmt"

func main(){
  radius,height:=getVal()
  volume:=cylinderVolume(radius,height)
  fmt.Println("Volume of cylinder: ",volume)
}

func getVal() (float64,float64){
  var radius float64
  var height float64
  fmt.Print("Enter radius:")
  fmt.Scanf("%f",&radius)
  fmt.Print("Enter height:")
  fmt.Scanf("%f",&height) //replace this line with the one commented below when running on Windows
  //fmt.Scanf("\n%f",&height)
  return radius,height
}

func cylinderVolume(radius float64, height float64) float64{
  const pi=3.14
  vol:=pi*radius*radius*height
  return vol
}
