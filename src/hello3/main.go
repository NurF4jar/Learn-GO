package main

import "fmt"
import "net/http"
import "html/template"
import "path"

func main()  {
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    // not yet implemente
    var filepath = path.Join("views", "index.html")
    var tmpl, err = template.ParseFiles(filepath)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    var data = map[string]interface{}{
        "title": "Learning Golang Web",
        "name":  "Nur Fajar",
    }

    err = tmpl.Execute(w, data)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)

    }

  })

  http.Handle("/static/",
    http.StripPrefix("/static/",
      http.FileServer(http.Dir("assets"))))

  fmt.Println("server started at 178.128.21.229:9000")
  http.ListenAndServe(":9000", nil)

}
