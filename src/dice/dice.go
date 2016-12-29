package main

import (
   "math/rand"
   "fmt"
   "time"
)

func randomize() {
   seconds := time.Now().Unix()
   fmt.Printf("seconds:%d\n", seconds)
   rand.Seed(seconds)
}

func rollDice() int {
   var n1, n2 int

   n1 = rand.Intn(6) + 1
   n2 = rand.Intn(6) + 1

   return n1 + n2
}

func main() {

   randomize()

   for i := 0; i<100000; i++ {
      x := rollDice()
   
      fmt.Printf("%6d x:%2d\n", i, x)

      if (x > 12) {
         break
      }
   }
}
