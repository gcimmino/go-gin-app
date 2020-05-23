package main

import "testing"

func TestGetAllArticles(test *testing.T)  {
  alist := getAllArticles()

  if len(alist) != len(articleList) {
    test.Fail()
  }

  for i, v := range alist {
    if v.Content != articleList[i].Content {
      v.ID != articleList[i].ID ||
      v.Title != articleList[i].Title {
        test.Failt()
        break
      }
    }
  }
}
