package main

func initializeRoutes()  {
  router.GET("/", showIndexPage)
}

func showIndexPage()  {
  router.GET("/", func(context *gin.Context)  {
    context.HTML(
      http.StatusOK,
      "index.html",
      gin.H{
        "title": "Home Page",
      },
    )
  })
}
