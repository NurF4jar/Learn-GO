package main

import "fmt"
import "net/http"

/*
func handlerIndex(w http.ResponseWriter, r *http.Request)  {
  var message = "Welcome"
  w.Write([]byte(message))
}

func handlerHello(w http.ResponseWriter, r *http.Request)  {
  var message = "Hello World!"
  w.Write([]byte(message))
}

func main()  {
  http.HandleFunc("/", handlerIndex)
  http.HandleFunc("/index", handlerIndex)
  http.HandleFunc("/hello", handlerHello)

  var address = "178.128.21.229:9000"
  fmt.Printf("server started at %s\n", address)
  server := new(http.Server)
  server.Addr = address
  err := server.ListenAndServe()
  if err != nil {
    fmt.Println(err.Error())
  }
*/

  // Routing
  // Penggunaan http.HandleFunc()
  func main()  {
    handlerIndex := func(w http.ResponseWriter, r *http.Request)  {
      w.Write([]byte("hello http.HandleFunc()"))
    }

    http.HandleFunc("/", handlerIndex)
    http.HandleFunc("index", handlerIndex)

    http.HandleFunc("/data", func(w http.ResponseWriter, r *http.Request)  {
      w.Write([]byte("Hello agan http.HandleFunc()"))
    })

    fmt.Println("Server started at 178.128.21.229:9000")
    http.ListenAndServe(":9000", nil)
  }
