package main

import "time"
import "fmt"
import "runtime"

func statusUpdate1() string {

   return "status1"
}

func statusUpdate2() string {

   return "status2"
}

func timer1() {

   c := time.Tick(309 * time.Millisecond)
   l := time.Now()

   for now := range c {
      d := now.Sub(l) / time.Microsecond
      l = now
      fmt.Printf("%v %s %d\n", now, statusUpdate1(), d)
   }
}

func timer2() {

   myClock := time.NewTicker(308571429 * time.Nanosecond)
   l := time.Now()
   
   for now := range myClock.C {
      d := now.Sub(l) / time.Microsecond
      l = now
      fmt.Printf("%v %s %d\n", now, statusUpdate2(), d)
   }
}

func main() {

   fmt.Printf("runtime.NumCPU() = %d\n", runtime.NumCPU())
   runtime.GOMAXPROCS(runtime.NumCPU())

   go timer1()
   go timer2()

   time.Sleep(10 * time.Second)
}
