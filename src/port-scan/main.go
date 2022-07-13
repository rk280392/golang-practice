package main

import (
   "fmt"
   "net"
   "sync"
)

func main(){
   var wg sync.WaitGroup //WaitGroup is a struct type and can be created as this 
   for i:=1; i<=1024;i++ {
      wg.Add(1)
      go func(j int) {
         defer wg.Done()
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
   wg.Wait()
}
