// Participatory Networking. Copyright (C) 2012-2013 Brown University
//
// Author: Andrew Ferguson <adf@cs.brown.edu>
// Author: Arjun Guha <arjun@cs.brown.edu>
//

package main

import (
  "runtime/pprof"
  "github.com/samuel/go-thrift/thrift"
  "goof/controller"

  "fmt"
  "log"
  "net"
  "net/rpc"
  "os"

  "pane"
)

func thrifttest(service pane.PaneService) {

  test := &pane.Principal {
    User: pane.ThriftString("root"),
  }

  fmt.Println(test);
}

func startPane() {
  pane_server := new(pane.PaneServer)
  pane_server.Init()
  rpc.RegisterName("Thrift", &pane.PaneServiceServer{pane_server})

  thrifttest(pane_server)

  ln, err := net.Listen("tcp", ":4242")
  if err != nil {
    panic(err)
  }

  for {
    conn, err := ln.Accept()
    if err != nil {
      fmt.Printf("ERROR: %+v\n", err)
      continue
    }
    fmt.Printf("New connection %+v\n", conn)
    go rpc.ServeCodec(thrift.NewServerCodec(conn,
                      thrift.NewBinaryProtocol(true, false)))
  }
}

func main() {
  f, _ := os.Create("profile")
  err2 := pprof.StartCPUProfile(f)
  if err2 != nil {
      panic(err2)
  }
  defer func() {
      log.Printf("Unprofiling")
  }()

  go startPane()

  log.Printf("Starting OpenFlow server ...")
  ctrl := controller.NewController()
  err := ctrl.Accept(6633, pane.NewSwitch)

  panic(err)
}
