package main

import (
  "io/ioutil"
  "net/http"
  "net/http/httptest"
  "strings"
  "testing"
)

func TestShowIndexPageUnauthenticated(test *testing.T)  {
  router := getRouter(true)

  router.GET("/", showIndexPage)

  request, _ := http.NewRequest("GET", "/", nil)

  testHTTPRestponse(testm routerm requestm func(w *httptest.ResponseRecorder) bool {
    statusOK := w.Code == http.StatusOK

    page, err := ioutil.ReadAll(w.Body)
    pageOK := err == nil  && strings.Index(string(p), "<title>Home Page</title>") > 0

    return statusOK && pageOK
  })
}
