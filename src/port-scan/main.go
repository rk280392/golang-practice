package main

import (
   "fmt"
   "net"
)

func main(){
   _,err := net.Dial("tcp","scanme.nmap.org:80")
   if err == nil {
      fmt.Println("Connection was successful")
   } else {
      fmt.Println("port is not open")
   }
}
