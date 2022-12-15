package bootstrap

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	log "github.com/sirupsen/logrus"
	"github.com/wechatgpt/wechatbot/config"
	"github.com/wechatgpt/wechatbot/handler/telegram"
	"github.com/wechatgpt/wechatbot/utils"
	"os"
	"strings"
	"time"
)

func StartTelegramBot() {
	telegramKey := os.Getenv("telegram")
	if len(telegramKey) == 0 {
		getConfig := config.GetConfig()
		if getConfig == nil {
			return
		}
		botConfig := getConfig.ChatGpt
		if botConfig.Telegram == nil {
			return
		}
		telegramKey = *botConfig.Telegram
		log.Info("读取本地本置文件中的telegram token:", telegramKey)
	} else {
		log.Info("找到环境变量: telegram token:", telegramKey)
	}
	bot, err := tgbotapi.NewBotAPI(telegramKey)
	if err != nil {
		return
	}

	bot.Debug = false
	log.Info("Authorized on account: ", bot.Self.UserName)
	u := tgbotapi.NewUpdate(0)

	updates := bot.GetUpdatesChan(u)
	time.Sleep(time.Millisecond * 500)
	for len(updates) != 0 {
		<-updates
	}

	for update := range updates {
		if update.Message == nil {
			continue
		}
		text := update.Message.Text
		chatID := update.Message.Chat.ID
		chatUserName := update.Message.Chat.UserName

		tgUserNameStr := os.Getenv("tg_whitelist")
		tgUserNames := strings.Split(tgUserNameStr, ",")
		if len(tgUserNames) > 0 && len(tgUserNameStr) > 0 {
			found := false
			for _, name := range tgUserNames {
				if name == chatUserName {
					found = true
					break
				}
			}

			if !found {
				log.Error("用户设置了私人私用，白名单以外的人不生效: ", chatUserName)
				continue
			}
		}

		tgKeyWord := os.Getenv("tg_keyword")
		var reply *string
		// 如果设置了关键字就以关键字为准，没设置就所有消息都监听
		if len(tgKeyWord) > 0 {
			content, key := utils.ContainsI(text, tgKeyWord)
			if len(key) == 0 {
				continue
			}
			splitItems := strings.Split(content, key)
			if len(splitItems) < 2 {
				continue
			}
			requestText := strings.TrimSpace(splitItems[1])
			log.Println("问题：", requestText)
			reply = telegram.Handle(requestText)
		} else {
			reply = telegram.Handle(text)
		}
		if reply == nil {
			continue
		}
		msg := tgbotapi.NewMessage(chatID, *reply)
		send, err := bot.Send(msg)
		if err != nil {
			log.Errorf("发送消息出错:%s", err.Error())
			continue
		}
		fmt.Println(send.Text)
	}
}
