package main

import (
   "github.com/codegangsta/martini"
   "net"
)

func main() {
  m := martini.Classic()
  interfaces, err := net.Interfaces()
  if err != nil {
    panic("Unable to get interfaces.")
  }
  mac := interfaces[1].HardwareAddr.String()

  m.Get("/", func() string {
    return "<h1>Welcome to Version 2<br><br>My MAC address is " + mac + "</h1>"
  })

  m.Run()
}
