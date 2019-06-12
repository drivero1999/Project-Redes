package main

import "fmt"
//goto Loop
func main() {
   /* local variable definition */
   var a int = 10
   Loop:
      fmt.Println("hola")

   if a==10{
      a++
      goto Loop
   }   
    
}
