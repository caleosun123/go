package main;

import (
  "fmt"
  "net/http"
)

func Handler(w http.ResponseWriter, r *http.request){
  server := New()
  server.get("/", func(context *Context){
    context.JSON(200,H{
      "message": "Hello, world!"
    })
  })
}
