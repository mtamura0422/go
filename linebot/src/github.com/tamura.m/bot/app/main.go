package main

import (

	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/tamura.m/bot/controllers"
	"github.com/BurntSushi/toml"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"os"
	"path/filepath"

)



func init() {

	InitGin()
}




func InitGin() {
	r := gin.Default()

	r.Use(initializer())
	r.GET("/", controllers.Top)
	r.POST("/hook", controllers.LineEntryPoint)
	r.GET("/hook", controllers.LineEntryPointTest)


	http.Handle("/", r)
}



// Midleware
func initializer() gin.HandlerFunc {
  return func(c *gin.Context) {

    ctx := appengine.NewContext(c.Request)

    exe, _ := os.Executable()
    log.Errorf(ctx, "%v", filepath.Dir(exe))

    var config controllers.Config

    _, err := toml.DecodeFile("config.toml", &config)
    if err != nil {
      log.Errorf(ctx, "toml file error!! %v", err)
    }

    c.Set("config", config)

  }
}
