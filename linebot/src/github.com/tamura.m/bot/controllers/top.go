package controllers


import (


  "net/http"
  "github.com/gin-gonic/gin"

//  "google.golang.org/appengine"
//  "google.golang.org/appengine/datastore"


)



// アクセスチェック用
func Test(c *gin.Context) {
  c.String(http.StatusOK, "Hello, world!!")
}


// アクセスチェック用
func Top(c *gin.Context) {
  c.String(http.StatusOK, "Hello, world!!")
}

