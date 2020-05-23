package main

import (
  "net/http"
  "github.com/gin-gonic/gin"
)

func showIndexPage(context *gin.Context)  {
  articles := getAllArticles()

  context.HTML(
    http.StatusOK,
    "index.html",
    gin.H{
      "title": "Home Page",
      "payload": articles,
    },
  )
}
