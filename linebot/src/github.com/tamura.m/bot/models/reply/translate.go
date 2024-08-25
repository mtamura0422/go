package reply


import (
  "net/http"
  "golang.org/x/net/context"

  "google.golang.org/appengine/log"
  "google.golang.org/api/translate/v2"

)


type Lang struct {
  FromLang, ToLang string
}


// 英語を翻訳
func (l Lang) TranslateText(g_client *http.Client, c context.Context, source string) (string, error) {

  t, err := translate.New(g_client)
  q := []string{source}
    ret, err := t.Translations.List(q, l.ToLang).Source(l.FromLang).Do()
    if err != nil {
        log.Errorf(c, "trans to create client: %v", err)
    }

  return ret.Translations[0].TranslatedText, err

}
