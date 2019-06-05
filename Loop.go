//Apprently there's only one kind of loop in Go and that's "for"!

package main
import "fmt"
func main(){

  fmt.Println("The first kind of for loop:")
  i:=0
  for i<10{  //coditional for
    fmt.Println(i)
    i++
  }

  fmt.Println("\nThe second kind of for loop:")
  for j:=0;j<10;j++{ //classic for
    fmt.Println(j)
  }

  fmt.Println("\nThe third kind of for loop:")
  k:=0
  for { //infinite for
    fmt.Println(k)
    if k<10{
      k++
    } else{
      break
    }
  }

}
