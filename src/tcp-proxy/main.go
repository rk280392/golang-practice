package main

import (
   "net"
   "log"
   "io"
)

func handler(src net.Conn) {
   dst,err:=net.Dial("tcp","learningdevops.com:80")
   if err != nil{
      log.Fatalln("Unable to connect to our unreachable host")
   }
   log.Println("Redirected to the domain")
   defer dst.Close()

   go func() {
      if _,err:=io.Copy(dst,src); err != nil {
         log.Fatalln(err)
      }
   }()
   if _,err:=io.Copy(src,dst); err !=nil {
      log.Fatalln(err)
   }
}

func main() {
   listener,err := net.Listen("tcp",":80")
   if err != nil {
      log.Fatalln("Unable to bind to port")
   }
   log.Println("Listening on port")

   for {
      conn,err := listener.Accept()
      if err !=nil {
         log.Fatalln("Unable to accept the connection")
      }
      log.Println("Connection accepted")
      go handler(conn)
   }
}

