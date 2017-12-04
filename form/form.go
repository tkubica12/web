package main

import (
   "github.com/go-martini/martini"
)

func main() {
  m := martini.Classic()

  page := `
  <h1>Azure + F5 rulez!<br><br>
  <form action="/send">
  <div class="container">
    <label><b>Username</b></label>
    <input type="text" placeholder="Enter Username" name="username" required>

    <label><b>Password</b></label>
    <input type="password" placeholder="Enter Password" name="password" required>

    <button type="submit">Login</button>
  </div>
</form>
  
  `

  m.Get("/", func() string {
    return page
  })

  m.Get("/send", func() string {
    return "Login processed"
  })

  m.RunOnAddr(":80")
}
