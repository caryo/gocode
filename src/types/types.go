package main

import "fmt"

/* be sure that the switch statement supports the type passed in */

func a(p interface{}) {

   switch p.(type) {
      case int:
         fmt.Printf("%d: Hello World.\n", p)
      case string:
         fmt.Printf("%s: Hello World.\n", p)
      default:
   }
}

/* in this example we call a function using two parameter types */

func main() {

   var i int = 1
   var s string = "Cary"
   var f float32 = 3.14

   a(i)
   a(s)
   a(f)
}
