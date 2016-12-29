package main

import "fmt"
import "log"
import "net/rpc"

type Arg struct {
   Increment int
}

type Result struct {
   Value int
}

func main() {
   client, err := rpc.Dial("tcp", "localhost:1234")
   if err != nil {
      log.Fatal("Value error:", err)
   }
   var r Result
   client.Call("GoCounter.Value", &Arg{1}, &r)
   fmt.Printf("%d\n", r.Value)
}
