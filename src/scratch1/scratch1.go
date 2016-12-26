package main

import (
   "fmt"
   "bufio"
   "os"
   "strconv"
   "strings"
)

func inputFloatValue(s string) float64 {
   var val float64
   var delim string

   reader := bufio.NewReader(os.Stdin)
   for {
      fmt.Printf("%s", s)
      str, err := reader.ReadString('\n')
      fmt.Printf("len:%d\n", len(str))
      idx := strings.Index(str,"\r")
      fmt.Printf("idx:%d\n", idx)
      if idx == -1 {
         delim = "\n"
      } else {
         delim = "\r"
      }
      if err != nil {
         fmt.Printf("input error - expected a floating point value\n")
      } else {
         var err2 error
	 str2 := strings.Split(str,delim)
         val, err2 = strconv.ParseFloat(str2[0],10)
	 if err2 != nil {
	    fmt.Printf("input error - expected a floating point value\n")
	 } else {
            break
	 }
      }
   }

   return val
}

func inputIntValue(s string) int64 {
   var val int64
   var delim string

   reader := bufio.NewReader(os.Stdin)
   for {
      fmt.Printf("%s", s)
      str, err := reader.ReadString('\n')
      fmt.Printf("len:%d\n", len(str))
      idx := strings.Index(str,"\r")
      fmt.Printf("idx:%d\n", idx)
      if idx == -1 {
         delim = "\n"
      } else {
         delim = "\r"
      }
      if err != nil {
         fmt.Printf("input error - expected an integer value\n")
      } else {
         var err2 error
	 str2 := strings.Split(str,delim)
         val, err2 = strconv.ParseInt(str2[0],10,64)
	 if err2 != nil {
	    fmt.Printf("input error - expected an integer value\n")
	    fmt.Printf("err2 = %v\n", err2)
	 } else {
	    if val > 9223372036854775807 {
	       fmt.Printf("input error - integer value exceeds max\n")
	    } else {
               break
	    }
	 }
      }
   }

   return val
}

func main() {

   var a int64
   var b int64
   var c int64
   var d float64
   var e float64
   var f float64

   a = inputIntValue("Enter an integer value for 'a' : ")
   b = inputIntValue("Enter an integer value for 'b' : ")

   c = a + b

   fmt.Printf("%d + %d = %d\n", a, b, c)

   d = inputFloatValue("Enter a floating point value for 'd' : ")
   e = inputFloatValue("Enter a floating point value for 'e' : ")

   f = d + e

   fmt.Printf("%f + %f = %f\n", d, e, f)
}
