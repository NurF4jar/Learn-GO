package main

import (
    "fmt"
    "github.com/labstack/echo"
    "net/http"
    "strings"
)

type M map[string]interface{}

func main() {
    r := echo.New()

    //method .STRING()
    // r.GET("/index", func(ctx echo.Context) error {
    //   data := "Hello from /index"
    //   return ctx.String(http.StatusOK, data)
    // })

    //method .HTML()
    r.GET("/html", func(ctx echo.Context) error {
      data := "Hello from /HTML"
      return ctx.HTML(http.StatusOK, data)
    })

    //method .HTML()
    r.GET("/index", func(ctx echo.Context) error {
      return ctx.Redirect(http.StatusTemporaryRedirect, "/json")
    })

    r.GET("/json", func(ctx echo.Context) error {
      data := M{"message": "Hello", "Counter": 2}
      return ctx.JSON(http.StatusOK, data)
    })

    // PARSING REQUEST
    // PARSING QUERY STRING
    r.GET("/page1", func(ctx echo.Context) error {
      name := ctx.QueryParam("name")
      data := fmt.Sprint("Hello %s", name)

      return ctx.String(http.StatusOK, data)

    })

    // PARSING URL PATH PARAM
    r.GET("/page2/:name", func(ctx echo.Context) error {
      name := ctx.Param("name")
      data := fmt.Sprintf("Hello %s", name)

      return ctx.String(http.StatusOK, data)

    } )

    // PARSING URL PATH PARAM DAN SETELAHNYA
    r.GET("/page3/:name/*", func(ctx echo.Context) error {
      name := ctx.Param("name")
      message := ctx.Param("*")

      data := fmt.Sprintf("Hello %s, I have message for you: %s", name, message)

      return ctx.String(http.StatusOK, data)

    })


    // PARSING FORM DATA
    r.POST("/page4", func(ctx echo.Context) error {
      name := ctx.FormValue("name")
      message := ctx.FormValue("message")

      data := fmt.Sprintf(
        "Hello %s, I have message for you: %s",
        name,
        strings.Replace(message, "/", "", 1),
      )

      return ctx.String(http.StatusOK, data)

    })

    r.Start(":9000")
    // fmt.Println("server started at :9000")
    // r.Start(":9000")
}
