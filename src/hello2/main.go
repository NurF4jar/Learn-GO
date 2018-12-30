package main

import "fmt"
import "net/http"

func main()  {
  http.Handle("/static/",
    http.StripPrefix("/static/",
      http.FileServer(http.Dir("assets"))))

  fmt.Println("Server started at 178.128.21.229:9000")
  http.ListenAndServe(":9000", nil)

}
