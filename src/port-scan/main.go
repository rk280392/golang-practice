package main

import (
   "fmt"
   "net"
)

func main(){
   for i:=1; i<=1024;i++ {
      go func (j int) {
         address := fmt.Sprintf("scanme.nmap.org:%d", j)
         conn,err := net.Dial("tcp",address)
         if err != nil {
            fmt.Printf("port %v is not open\n",j)
            return
         } else {
            conn.Close()
            fmt.Printf("Connection was successful on port %v\n",j)
         }
      }(i)
   }
}
