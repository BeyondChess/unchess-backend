package main

import (
  "apps/auth/internal/server"
  "fmt"
)

func main() {

  server := server.NewServer()

  err := server.ListenAndServe()
  if err != nil {
    panic(fmt.Sprintf("cannot start server: %s", err))
  }
  
}
