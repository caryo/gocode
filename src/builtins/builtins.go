package main

import (
  "fmt"
  "flag"
)

func countToN(n *int) {

   for i := 0; i<*n; i++ {
      fmt.Printf("%d ", i+1)
   }
   fmt.Printf("\n")
}

func fillToN(a []int, n *int) {

   for i := 0; i<*n; i++ {
      a[i] = i + 1
   }
}

func printArray(a []int, n *int) {

   l := len(a)
   fmt.Printf("len(a) = %d\n", l)

   c := cap(a)
   fmt.Printf("cap(a) = %d\n", c)

   ub := *n

/*
   if ub > len(a) {
      ub = len(a)
   }
*/

   for i := 0; i<ub; i++ {
      fmt.Printf("%d ", a[i])
   }
   fmt.Printf("\n")
}

func cleanup() {
   fmt.Printf("\n")
   if x := recover(); x != nil {
      fmt.Printf("run time panic: %v\n", x)
   }
   fmt.Printf("Done.\n")
}

func main() {

   defer cleanup()

   _n := flag.Int("n",10,"number of items to allocate")
   _s := flag.Int("s",20,"max size of item array")

   flag.Parse()

   fmt.Printf("Hello World.\n")

   n := new(int)
   s := new(int)

   n = _n
   s = _s

   countToN(n)

   a := make([]int, *n, *s)

   printArray(a, n)
   printArray(a, s)

   fillToN(a, n)

   printArray(a, n)
   printArray(a, s)
}
