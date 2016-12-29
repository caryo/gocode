package main

import (
   "bytes"
   "fmt"
   "os"
   "log"
)

func main() {

   var buf bytes.Buffer

   fmt.Printf("Address Book:\n")

   logger := log.New(&buf, "logger: ", log.Lshortfile)
   logger.Print("opening addrbook.dat")

   _, err := os.Open("addrbook.dat")
   if  err != nil {
      if !os.IsNotExist(err) {
         fmt.Println(err)
         logger.Print(err)
         fmt.Print(&buf)
         fmt.Print()
         log.Fatal(err)
      } else {
         log.Print(err)
         logger.Print(err)
      }
   }
   fmt.Printf("ready\n")
   fmt.Print(&buf)
}
