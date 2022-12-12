package bootstrap

import (
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"log"
)

// StartLineBot https://manager.line.biz/account/@383wbcrn/setting/messaging-api
func StartLineBot() {
	bot, err := linebot.New(
		"",
		"")

	// エラーに値があればログに出力し終了する
	if err != nil {
		log.Fatal(err)
	}
	// weatherパッケージパッケージから天気情報の文字列をを取得する
	result := "hello"
	// エラーに値があればログに出力し終了する
	if err != nil {
		log.Fatal(err)
	}
	// テキストメッセージを生成する
	message := linebot.NewTextMessage(result)
	// テキストメッセージを友達登録しているユーザー全員に配信する
	if _, err := bot.BroadcastMessage(message).Do(); err != nil {
		log.Fatal(err)
	}
}
