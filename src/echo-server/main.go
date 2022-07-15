package main

import (
   "net"
   "log"
   "bufio"
)

func echo(conn net.Conn) {
   defer conn.Close()

   reader:=bufio.NewReader(conn)
   s,err:=reader.ReadString('\n')
   if err != nil {
      log.Fatalln("Unable to read data")
   }

   log.Printf("Read %d bytes :%s", len(s),s)

   log.Print("Writing data")
   writer := bufio.NewWriter(conn)
   if _,err:=writer.WriteString(s); err != nil {
      log.Fatalln("Unable to write data")
   }
   writer.Flush()
}

func main() {
   port := 20080
   listner,err:=net.Listen("tcp",":20080")
   if err != nil {
      log.Fatalln("Unable to bind port")
   }
   log.Printf("Listening on 0.0.0.0:%d", port)

   for {
      conn,err:=listner.Accept()
      log.Println("Received connection")
      if err != nil {
         log.Fatalln("Unable to accept connection")
      }
      go echo(conn)
   }
}
