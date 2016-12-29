package main

import "fmt"

func noresult() {
   panic("noresult")
   return
}

func continueYN() byte {
   var c byte
   fmt.Printf("Continue (y/n) ")
   fmt.Scanf("%c", &c)
   fmt.Printf("%c\n", c)
   return c
}

func countToN(n *int) {
   for i := 0; i < *n; i++ {
      fmt.Printf("%d ", i+1)
   }
   fmt.Printf("\n")
   *n = 0
}

func main() {
   a := 1
   if a == 1 {
      c := continueYN()
      if c == 'n' {
         noresult()
      }
   }
   var n int = 10
   countToN(&n)
   fmt.Printf("n=%d\n", n)
   fmt.Printf("%d\n", a)
   return

}
