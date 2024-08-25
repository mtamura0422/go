package controllers


import (


  "net/http"
  "github.com/gin-gonic/gin"
  "github.com/line/line-bot-sdk-go/linebot"
  "google.golang.org/appengine"
//  "google.golang.org/appengine/datastore"
  "google.golang.org/appengine/log"
  "strconv"
  "bot/models/reply"
  "bot/models/google"

)



type Config struct {
  Api APIList
}


type APIList struct {
  A3rt A3rtList
  Line LineConfig
  Google GoogleList
  Weather WeatherList
}

type A3rtList struct {
  Talk A3rtTalkList
}


type A3rtTalkList struct {
  Key string
}

type WeatherList struct {
  Key string
}


type LineConfig struct {
  Secret string
  Token string
  Key string
}



type GoogleList struct {
    Translate GoogleTranslateList
}


type GoogleTranslateList struct {
    Key string
}




const MyStructKind string = "ReplayMesssage"
const MessageFile string = "files/message.csv"




type ReplayMesssage struct {
//  Id string `datastore:"-" goon:"id"`
  ReplayType int `datastore:"replay_type"`
  Message string `datastore:"message,noindex"`
}


type User struct {
  Id   int `datastore:"-" goon:"id"`
  UserId string `datastore:"noindex"`
  Name string `datastore:"noindex"`
}

// LineBot エントリーポイント
func LineEntryPointTest(c *gin.Context) {
  ctx := appengine.NewContext(c.Request)
  config := c.MustGet("config").(Config)
  log.Errorf(ctx, "%v", config)
 c.String(http.StatusOK, "Hello, world!!")
}


// LineBot エントリーポイント
func LineEntryPoint(c *gin.Context) {

  var replyMsg string = ""

  ctx := appengine.NewContext(c.Request)
  config := c.MustGet("config").(Config)
log.Errorf(ctx, "%v", "tuuka")
  g_client := google.New(ctx, config.Api.Google.Translate.Key)
  // linbotオブジェクト生成(clientはappengineでfetchしたものを使用する)
  bot, err := linebot.New(config.Api.Line.Secret, config.Api.Line.Token, linebot.WithHTTPClient(g_client))
  if err != nil {
    log.Errorf(ctx, "%v", err)
  } else {
  log.Errorf(ctx, "%v", bot)
    // リクエストをパース
    events, err := bot.ParseRequest(c.Request)
    for _, event := range events {
      if event.Type == linebot.EventTypeMessage {

        switch message := event.Message.(type) {
        // テキストメッセージの場合
        case *linebot.TextMessage:

          replyMsg = reply.TextReply(ctx, config.Api.A3rt.Talk.Key, event.Source.UserID, bot, message, g_client)

        // 画像の場合
        case *linebot.ImageMessage:
          replyMsg = reply.ImageReply(g_client, ctx, bot, message)

        // 位置情報の場合
        case *linebot.LocationMessage:

//          var messages []linebot.SendingMessage
          hinanInfo := reply.LocationReply(ctx, config.Api.Weather.Key, event.Source.UserID, bot, message)


      //    if _, err = bot.ReplyMessage(event.ReplyToken, messages...).WithContext(ctx).Do(); err != nil {
      //      log.Errorf(ctx, "%v", err)
      //    }


          // 緯度経度を文字列に変換する
          hinanIdo, _ := strconv.ParseFloat(hinanInfo[14], 64)
          hinanKdo, _ := strconv.ParseFloat(hinanInfo[15], 64)


          if _, err = bot.ReplyMessage(event.ReplyToken,linebot.NewLocationMessage(
              hinanInfo[3],
              hinanInfo[4],
              hinanIdo,
              hinanKdo)).Do(); err != nil {
            log.Errorf(ctx, "%v", hinanInfo)
          }
        }

        if replyMsg != "" {
          if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(replyMsg)).WithContext(ctx).Do(); err != nil {
            log.Errorf(ctx, "%v", err)
          }
        }
      }
    }
  }
}
