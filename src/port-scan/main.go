package main

import (
   "fmt"
   "net"
   "sync"
)

func main(){
   var wg sync.WaitGroup //WaitGroup is a struct type and acts as a synchronized counter 
   for i:=1; i<=1024;i++ {
      wg.Add(1) // increment counter via wg.Add(1) each time a goroutine is created to scan a port
      go func(j int) {
         defer wg.Done() // decrements the counter whenever one unit of work has been performed 
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
   // main() function calls wg.Wait(), which blocks until all the work has been done and your counter has returned to zero
   wg.Wait()
}
