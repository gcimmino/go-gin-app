package main

import (
  "strconv"
  "net/http"
  "github.com/gin-gonic/gin"
)

func showIndexPage(context *gin.Context)  {
  articles := getAllArticles()

  render(context, gin.H{
    "title": "Home Page",
    "payload": articles,
  }, "index.html")
}

func getArticle(context *gin.Context)  {
  if articleID, err := strconv.Atoi(context.Param("article_id")); err == nil {
    if article, err := getArticleByID(articleID); err == nil {
      render(context, gin.H{
        "title": article.Title,
        "payload": article,
      }, "article.html")
    } else {
      context.AbortWithError(http.StatusNotFound, err)
    }
  } else {
    context.AbortWithStatus(http.StatusNotFound)
  }
}
