package main

import (
  "strconv"
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

func getArticle(context *gin.Context)  {
  if articleID, err := strconv.Atoi(context.Param("article_id")); err == nil {
    if article, err := getArticleByID(articleID); err == nil {
      context.HTML(
        http.StatusOK,
        "article.html",
        gin.H{
          "title": article.Title,
          "payload": article,
        },
      )
    } else {
      context.AbortWithError(http.StatusNotFound, err)
    }
  } else {
    context.AbortWithStatus(http.StatusNotFound)
  }
}
