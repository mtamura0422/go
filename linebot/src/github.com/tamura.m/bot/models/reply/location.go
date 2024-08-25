package reply

import (
//  "fmt"
  "golang.org/x/net/context"
  "github.com/line/line-bot-sdk-go/linebot"
  "google.golang.org/appengine/log"
//  "net/http"
  "strconv"
//  "encoding/json"
//  "io/ioutil"
//  "google.golang.org/appengine/urlfetch"
  "encoding/csv"
  "os"
  "io"
  "math"
  "reflect"
)


type Wdata struct {
  Weather []Weather `json:"weather"`
  Info    Info      `json:"main"`
}


// Weather represents weather item.
type Weather struct {
  Main string `json:"main"`
  Icon string `json:"icon"`
}

// Info represents main item.
type Info struct {
  Temp     float32 `json:"temp"`     // 気温(ケルビン)
  Humidity float32 `json:"humidity"` // 湿度
}





//func LocationReply(c context.Context, weather_api_key string, user_id string, bot *linebot.Client, message *linebot.LocationMessage) []linebot.SendingMessage {
func LocationReply(c context.Context, weather_api_key string, user_id string, bot *linebot.Client, message *linebot.LocationMessage) []string {


/*
  var replyMsg []linebot.SendingMessage
  log.Errorf(c, "%v", message)
  // 経度緯度取得
//  location, _ := c.LocationContent()
  lat := strconv.FormatFloat(message.Latitude, 'f', 6, 64)
  lon := strconv.FormatFloat(message.Longitude, 'f', 6, 64)


  url := "http://api.openweathermap.org/data/2.5/weather?lat=" + lat + "&lon=" + lon + "&APPID=" + weather_api_key
  httpClient := urlfetch.Client(c)

  resp, err := httpClient.Get(url)
  if err != nil {
    log.Errorf(c, "%v", err)
  }
  if resp != nil {
      defer resp.Body.Close()
  }

  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    log.Errorf(c, "%v", err)
  }

log.Errorf(c, "%v", "tuuka2")
  jsonBytes := ([]byte)(string(body[:]))
  wdata := new(Wdata)
  if err := json.Unmarshal(jsonBytes, wdata); err != nil {
      fmt.Println("JSON Unmarshal error:", err)
      return replyMsg
  }
log.Errorf(c, "%v", "tuuka3")


  replyMsg = append(replyMsg, linebot.NewTextMessage("現在の天気をお知らせします。"))
  replyMsg = append(replyMsg, linebot.NewTextMessage("天気 : "+wdata.Weather[0].Main))
  replyMsg = append(replyMsg, linebot.NewTextMessage("気温 : "+fmt.Sprintf("%.2f", (wdata.Info.Temp-273.15))))
  replyMsg = append(replyMsg, linebot.NewTextMessage("湿度 : "+fmt.Sprintf("%.2f", wdata.Info.Humidity)))


  imgUrl := "https://openweathermap.org/img/w/"+wdata.Weather[0].Icon+".png"
  replyMsg = append(replyMsg, linebot.NewImageMessage(imgUrl, imgUrl))


*/
log.Errorf(c, "%v", "tuuka4")

  // 近くの避難所を取得する
  hinanInfo := loadHinanMap(c, message.Latitude, message.Longitude)

log.Errorf(c, "%v", "tuuka5")
  return hinanInfo

}





func distance(ido1 float64, kdo1 float64, ido2 float64, kdo2 float64) float64 {

  rlat1 := ido1 * math.Pi / 180
  rlng1 := kdo1 * math.Pi / 180
  rlat2 := ido2 * math.Pi / 180
  rlng2 := kdo2 * math.Pi / 180


  a := math.Sin(rlat1) * math.Sin(rlat2) +
      math.Cos(rlat1) * math.Cos(rlat2) *
      math.Cos(rlng1 - rlng2)

  rr := math.Acos(a)

  earth_radius :=  6378140.0
  distance := earth_radius * rr

  return distance

}



func loadHinanMap(c context.Context, ido float64, kdo float64) []string {

  var err error
  fp, err := os.Open("blame.csv")
  if err != nil {
    log.Errorf(c, "%v", err)
    panic(err)
  }
  defer fp.Close()


  // 現在地付近の場所に絞り込む
  ido_from := ido - 0.05
  ido_to := ido + 0.05
  kdo_from := kdo - 0.05
  kdo_to := kdo + 0.05

  min_dist := 0.00

  var min_row []string


  reader := csv.NewReader(fp)
  reader.Comma = ','
  reader.LazyQuotes = true // ダブルクオートを厳密にチェックしない！
  for {
    record, err := reader.Read()
    if err == io.EOF {
          break
    } else if err != nil {
          log.Errorf(c, "%v", err)
    }
   // log.Errorf(c, "%v", record)

hinan_ido, _ := strconv.ParseFloat(record[14], 64)
hinan_kdo, _ := strconv.ParseFloat(record[15], 64)

    if (ido_from <= hinan_ido && hinan_ido <= ido_to &&
        kdo_from <= hinan_kdo && hinan_kdo <= kdo_to) {

      dist := distance(ido, kdo, hinan_ido, hinan_kdo)


log.Errorf(c, "dist %v", dist)
log.Errorf(c, "record %v", record)
log.Errorf(c, "record type %v", reflect.TypeOf(record))



      if (min_dist == 0.00  || dist < min_dist) {
        min_dist = dist
        min_row = record
      }
log.Errorf(c, "min_dist %v", min_dist)
log.Errorf(c, "min_row %v", min_row[0])


    }
  }

  return min_row

}
