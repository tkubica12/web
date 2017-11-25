package main

import (
   "github.com/go-martini/martini"
   "os/exec"
)

func main() {
  m := martini.Classic()

  uuid, err := exec.Command("uuidgen").Output()
  if err != nil {
    panic("Unable to generate uuid")
  }

  m.Get("/", func() string {
    return "<h1>Azure + F5 rulez!<br><br>This is server " + string(uuid) + "</h1>"
  })

  m.RunOnAddr(":80")
}
