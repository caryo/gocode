package main

import (
   "fmt"
   "time"
)

func listen(c <-chan int) {

   var val int
   var status bool

   fmt.Printf("listening on channel ...\n")

   status = true

   for status {
      val, status = <-c
      fmt.Printf("val=%d, status=%v\n", val, status)
   }

   fmt.Printf("no longer listening on channel\n")

}

func main() {

   var c chan int

   fmt.Printf("Hello World.\n")

   c = make(chan int)

   go listen((<-chan int)(c))

   for x := 2; x > 0; x-- {
      time.Sleep(1 * time.Second)
      fmt.Printf("%d ", x)
   }
   fmt.Printf("\n")

   fmt.Printf("sending 1 on channel\n")

   s := (chan<- int)(c)

   s <- 1

   for x := 2; x > 0; x-- {
      time.Sleep(1 * time.Second)
      fmt.Printf("%d ", x)
   }
   fmt.Printf("\n")

   fmt.Printf("closing channel\n")
   close(s)

   for x := 2; x > 0; x-- {
      time.Sleep(1 * time.Second)
      fmt.Printf("%d ", x)
   }
   fmt.Printf("\n")
}
