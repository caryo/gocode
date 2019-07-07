/*
 * http://people.math.sfu.ca/~cbm/aands/page_297.htm
 * 7.1.6 Series expansion
 * https://math.stackexchange.com/questions/1776998/approximating-the-normal-cdf-for-0-leq-x-leq-7
 */
package main

import (
   "fmt"
   "math"
)

func fact(x int) uint64 {
   var val uint64
   if x == 0 {
      val = uint64(1)
   } else {
      val = uint64(x)
   }
   for i := 1 ; i < x-1; i++ {
      val *= uint64(x - i);
   }
   return val
}

func erf(z float64) float64 {
   val := 2.0 / math.Sqrt( math.Pi )
   a := 0.
   for n := 0.; n < 25.; n += 1.0 {
      a += ( ( math.Pow( -1.0, n ) * math.Pow( z, (2.0 * n + 1) ) ) /
             ( float64( fact( int(n) ) ) * (2.0 * n + 1) ) )
   }
   val *= a
   return val
}

func main() {
   var z, z0, z1 float64
   fmt.Printf("%6s %8s %8s %8s %8s %8s %8s %8s %8s %8s %8s\n",
      "x:", ".00", ".01", ".02", ".03", ".04", ".05", ".06", ".07", ".08", ".09")
   for z0 = -4.0; z0 <= 4.0; z0 += .1 {
      fmt.Printf("%5.1f: ", z0)
      for z1 = .00; z1 <= .09; z1 += .01 {
         z = z0 + z1
         ans := .5 * (1 + erf(z/math.Sqrt(2)))
         if ans < 0. {
            ans = 0.
         } else if ans > 1 {
            ans = 1.
         }
         fmt.Printf("%8.5f ", ans)
      }
      fmt.Println()
   }
   fmt.Println();
   z = -3.0
   fmt.Printf("z: %8.5f, erf(%8.5f): %8.5f\n", z, z, .5 * (1 + erf(z/math.Sqrt(2))) )
   z = -2.0
   fmt.Printf("z: %8.5f, erf(%8.5f): %8.5f\n", z, z, .5 * (1 + erf(z/math.Sqrt(2))) )
   z = -1.0
   fmt.Printf("z: %8.5f, erf(%8.5f): %8.5f\n", z, z, .5 * (1 + erf(z/math.Sqrt(2))) )
   z = .0
   fmt.Printf("z: %8.5f, erf(%8.5f): %8.5f\n", z, z, .5 * (1 + erf(z/math.Sqrt(2))) )
   z = 1.0
   fmt.Printf("z: %8.5f, erf(%8.5f): %8.5f\n", z, z, .5 * (1 + erf(z/math.Sqrt(2))) )
   z = 2.0
   fmt.Printf("z: %8.5f, erf(%8.5f): %8.5f\n", z, z, .5 * (1 + erf(z/math.Sqrt(2))) )
   z = 3.0
   fmt.Printf("z: %8.5f, erf(%8.5f): %8.5f\n", z, z, .5 * (1 + erf(z/math.Sqrt(2))) )
}
