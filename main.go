package main

import (
  "github.com/gin-gonic/gin"
  "net/http"
)

var router *gin.Engine

func main() {
  router = gin.Default()

  router.LoadHTMLGlob("templates/*")

  initializeRoutes()

  router.Run()
}

func render(context *gin.Context, data gin.H, templateName string)  {
  switch context.Request.Header.Get("Accept") {
  case "application/json":
    context.JSON(http.StatusOK, data["payload"])
  case "application/xml":
    context.XML(http.StatusOK, data["payload"])
  default:
    context.HTML(http.StatusOK, templateName, data)
  }
}
