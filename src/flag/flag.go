package main

import "fmt"
import "flag"

func main() {

   src := flag.String("s", "source", "source IP address/hostname")
   dst := flag.String("d", "destination", "destination IP address/hostname")

   flag.Parse()

   fmt.Printf("src=%s, dst=%s\n", *src, *dst)
   fmt.Printf("Hello World.\n")
}
