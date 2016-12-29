package main

import "net"
import "net/rpc"

type Counter struct {
   count int
}

type Arg struct {
   Increment int
}


type Result struct {
   Value int
}

func (c *Counter) Value(in *Arg, out *Result) error {
   c.count += in.Increment
   out.Value = c.count
   return nil
}

func main() {
   server := rpc.NewServer()
   server.RegisterName("GoCounter", new(Counter))
   l, _ := net.Listen("tcp", "localhost:1234")
   server.Accept(l)
}
