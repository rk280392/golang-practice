package main

import (
   "fmt"
   "net"
)

func main(){
   for i:=1; i<=1024;i++ {
      address := fmt.Sprintf("scanme.nmap.org:%d", i)
      conn,err := net.Dial("tcp",address)
      if err != nil {
         fmt.Printf("port %v is not open\n",i)
         continue
      } else {
         conn.Close()
         fmt.Printf("Connection was successful on port %v\n",i)
      }
   }
}
