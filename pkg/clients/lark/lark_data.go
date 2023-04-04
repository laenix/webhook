package lark

import "time"

func GetLarkRequestData(title, repo string) string {

	reqdata := `{
	"msg_type": "interactive",
	"card": {
		"config": {
			"wide_screen_mode": true
		},
		"header": {
			"title": {
				"tag": "plain_text",
				"content": 🟢"` + title + `"
			},
			"template": "blue"
		},
		"elements": [{
				"tag": "div",
				"text": {
					"content": "⚙ 目标仓库:  ` + repo + `  \n📅 推送时间: ` + time.Now().Local().String() + `",
					"tag": "lark_md"
				}
			},
		]}
	}`
	return reqdata
}
