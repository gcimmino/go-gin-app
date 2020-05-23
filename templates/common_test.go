package main

import (
  "net/http"
  "net/http/httptest"
  "os"
  "testing"
  "github.com/gin-gonic/gin"
)

var tmpArticleList []article

// Setup function executed before test functions
func TestMain(main *testing.M)  {
  gin.SetMode(gin.TestMode)

  os.Exit(main.Run())
}

// Helper to create router
func getRouter(withTemplates bool) *gin.Engine  {
  router := gin.Default()
  if withTemplates {
    router.LoadHTMLGlobal("templates/*")
  }

  return router
}

// Helper to process request and test response
func testHTTPRestponse(test *testing.T, router *gin.Engine, request *http.Request, f func(w *httptest.ResponseRecorder) bool)  {
  // Response recorder
  w := httptest.NetRecorder()

  // Create service and process the request
  router.ServeHTTP(w, request)

  if !f(w) {
    test.Fail()
  }
}


// Store main lists into temporary one
func saveLists()  {
  tmpArticleList = articleList
}

// Restore the main lists
func restoreLists()  {
  articleList = tmpArticleList
}
