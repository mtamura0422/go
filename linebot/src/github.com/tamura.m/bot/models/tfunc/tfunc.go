package tfunc


import (
//  "fmt"
  "time"
)

func FormatAsDbDate(t time.Time) string {

  return t.Format("20060102150405")

}



func FormatAsViewDate(t time.Time) string {

  return t.Format("2006/01/02 15:04:05")

}




func FormatAsOrderStatus(status int) string {

  var ret string

  switch status {
  case 0: ret = "未対応"
  case 1: ret = "完了"
  case 2: ret = "キャンセル"
  }

  return ret

}
