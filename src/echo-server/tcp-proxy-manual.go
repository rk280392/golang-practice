package main

import (
   "log"
   "net"
   "io"
)

func echo(conn net.Conn) {
   defer conn.Close()

   // buffer to store data
   b:=make([]byte,512)

   for {
      size,err:=conn.Read(b[0:])
      if err == io.EOF{
         log.Println("Client disconnected")
         break
      }

      if err != nil {
         log.Fatalln("Unexpected error")
         log.Println("Unexpected error")
         break
      }

      log.Printf("Received %d bytes: %s\n", size, string(b))

      log.Println("Writing data")
      if _,err:=conn.Write(b[0:size]); err !=nil {
         log.Fatalln("Unable to write data")
      }
   }
}

func main() {
   port := 20080
   listner,err := net.Listen("tcp",":20080")
   if err != nil {
      log.Fatalf("Unable to bind to port %d", port)
   }
   log.Printf("Listening on 0.0.0.0:%d", port)

   for {
      conn,err := listner.Accept()
      log.Println("Received connection")
      if err != nil {
         log.Fatalln("Unable to accept connection")
      }
      go echo(conn)
   }
}

