package reply

import (
	"fmt"
	"regexp"
	"golang.org/x/net/context"
	"github.com/line/line-bot-sdk-go/linebot"
	"google.golang.org/appengine/log"
	"net/http"
	"github.com/tamura.m/bot/models/talk"
)


const TRANSLATE_HEADER = "^:訳して"

func TextReply(c context.Context, talk_api_key string, user_id string, bot *linebot.Client, message  *linebot.TextMessage, g_client *http.Client) string {

	var replyMsg string
	var user_name string = "名無し"

	if user_id != "" {
	//	user_profile, err := bot.GetProfile(event.Source.UserID).Do()

		user_profile, err := bot.GetProfile(user_id).Do()
		if err == nil {
			user_name = user_profile.DisplayName
		}
	}

	ret_str := regexp.MustCompile(TRANSLATE_HEADER).ReplaceAllString(message.Text, "")

	if ret_str != message.Text {
		la := Lang{"ja", "en"}
		ret_str, err := la.TranslateText(g_client, c, ret_str)
		if err != nil {
			log.Errorf(c, "translateText error:%v", err)
		}

		replyMsg = fmt.Sprintf("%sちゃん\n「%s」\nだよ!!", user_name, ret_str)
	} else {

		// linbotオブジェクト生成(clientはappengineでfetchしたものを使用する)
		t, err := talk.New(talk.WithHTTPClient(g_client))
		if err != nil {
			log.Errorf(c, "%v", err)
		}

		ret, err := t.ReplyMessage(talk_api_key, message.Text).WithContext(c).Do()

		if err != nil {
			log.Errorf(c, "%v", err)
		}

		replyMsg = ret.Results[0].Reply

	}

	return replyMsg
}
