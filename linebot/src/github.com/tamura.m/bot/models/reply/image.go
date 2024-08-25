package reply

import ( 
	"fmt"
	"math/rand"
	"time"
	"net/http"
	"io/ioutil"
	"encoding/base64"

	"github.com/line/line-bot-sdk-go/linebot"

	"golang.org/x/net/context"
	"google.golang.org/api/vision/v1"
	"google.golang.org/appengine/log"

)




func ImageReply (g_client *http.Client, c context.Context, bot *linebot.Client, message  *linebot.ImageMessage) string {

	var replyMsg = "なんの写真？"

	// 画像を取得
	content, err := bot.GetMessageContent(message.ID).WithContext(c).Do()
	if err != nil {
		log.Errorf(c, "get content:%v", err)
	}

	// 画像読み込み
	tmp_img, err := ioutil.ReadAll(content.Content)
	if err != nil {
		log.Errorf(c, "read image:%v", err)
	}


	// 1つの画像に対して複数のFeatureを指定できる
	req := &vision.AnnotateImageRequest{
		Image: &vision.Image{Content: base64.StdEncoding.EncodeToString(tmp_img)},
		Features: []*vision.Feature{
			&vision.Feature{
				Type:       "LABEL_DETECTION",
				MaxResults: 10,
			},
		},
	}

	// 1回の呼び出しで複数の処理を要求できる
	batch := &vision.BatchAnnotateImagesRequest{
		Requests: []*vision.AnnotateImageRequest{req},
	}

	
	svc, err := vision.New(g_client)
	if err != nil {
		log.Errorf(c, "vision new:%v", err)
	}

	res, err := svc.Images.Annotate(batch).Do()
	if err != nil {
		log.Errorf(c, "annotate:%v", err)
	}

	var descriptions []string

	for _, ret := range res.Responses[0].LabelAnnotations {
		log.Infof(c, "ret %v", ret)
		// 結果を取得
		descriptions = append(descriptions, ret.Description)
	}


	if len(descriptions) > 0 {
		rand.Seed(time.Now().UnixNano())
		i := rand.Intn(len(descriptions))

		ret := descriptions[i]
	
		la := Lang{"en", "ja"}
		ret_t, err := la.TranslateText(g_client, c, ret) 
		if err != nil {
			log.Errorf(c, "translateText error:%v", err)
		}

		log.Infof(c, "trans %v", ret_t)
		replyMsg = fmt.Sprintf("%sじゃーん!!", ret_t) 
	}

	return replyMsg
	
}


