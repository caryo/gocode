package main

import "fmt"
import "math"

func main() {

   var p,i,r,a,pp,ip float64
   var n int32
   var c byte

   fmt.Printf("Enter principal amount: ")
   fmt.Scanf("%f\n", &p)

   fmt.Printf("Enter interest rate: ")
   fmt.Scanf("%f\n", &i)

   fmt.Printf("Enter number of payments: ")
   fmt.Scanf("%d\n", &n)

   fmt.Printf("Principal amount: %f\n", p)
   fmt.Printf("Interest rate: %f\n", i)
   fmt.Printf("Number of payments: %d\n", n)

   r = 1.0 + i / float64(12.0)

   a = p * ((i/(float64(12.0)))*math.Pow(r,float64(n)) / (math.Pow(r,float64(n))-1.0))

   fmt.Printf("Payment amount: %8.2f\n", a)

   fmt.Printf("Print Schedule? (y/n) ")
   fmt.Scanf("%c\n", &c)

   if c == 'y' {
      fmt.Printf("Schedule:\n")
      fmt.Printf("Payment Principal Interest Amount Remaining\n")
      k := 1
      for j := int32(1); j <= n; j++ {
         pp = p - (p * r - a )
	 ip = a - pp
         p = p * r - a;
	 // fmt.Printf("%d %f %f %f\n", j, pp, ip, p)
	 fmt.Printf("%8d %8.2f %8.2f %8.2f %8.2f\n", j, pp, ip, (pp + ip), p)
	 if k < 12 {
	    k = k + 1
	 } else {
	    fmt.Printf("Press any key to continue ('q' to quit)")
	    c = 0
	    fmt.Scanf("%c\n", &c)
	    fmt.Printf("\n")
	    if c == 'q' {
	       break
	    }
	    k = 1
	 }
      }
   }
}
