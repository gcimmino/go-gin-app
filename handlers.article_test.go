package main

import (
  "io/ioutil"
  "net/http"
  "net/http/httptest"
  "strings"
  "testing"
)

func TestShowIndexPageUnauthenticated(test *testing.T) {
  router := getRouter(true)

  router.GET("/", showIndexPage)

  request, _ := http.NewRequest("GET", "/", nil)

  testHTTPResponse(test, router, request, func(w *httptest.ResponseRecorder) bool {
    statusOK := w.Code == http.StatusOK

    page, err := ioutil.ReadAll(w.Body)
    pageOK := err == nil  && strings.Index(string(page), "<title>Home Page</title>") > 0

    return statusOK && pageOK
  })
}

func TestShowIndexPageAuthenticated(test *testing.T)  {
  w := httptest.NewRecorder()

  router := getRouter(true)

  http.SetCookie(w, &http.Cookie{Name: "token", Value: "123"})

  router.GET("/", showIndexPage)

  request, _ := http.NewRequest("GET", "/", nil)
  request.Header = http.Header{"Cookie": w.HeaderMap["Set-Cookie"]}

  router.ServeHTTP(w, request)

  if w.Code != http.StatusOK {
    test.Fail()
  }

  page, err := ioutil.ReadAll(w.Body)

  if err != nil || strings.Index(string(page), "<title>Home Page</title>") < 0 {
    test.Fail()
  }
}

func TestArticleUnauthenticated(test *testing.T)  {
  router := getRouter(true)
  router.GET("/article/view/:article_id", getArticle)
  request, _ := http.NewRequest("GET", "/article/view/1", nil)

  testHTTPResponse(test, router, request, func(w *httptest.ResponseRecorder) bool {
    statusOK := w.Code == http.StatusOK

    page, err := ioutil.ReadAll(w.Body)
    pageOK := err == nil && strings.Index(string(page), "<title>Article 1</title>") > 0

    return statusOK && pageOK
  })
}
