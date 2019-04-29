package bot_handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/linebot"
	"log"
	"os"
	"strings"
)

func SampleBotHandler(c *gin.Context) {
	log.Printf("context = %+v", c)
	bot, err := linebot.New(
		os.Getenv("Secret"),
		os.Getenv("Token"))
	if err != nil {
		log.Fatalf("New client error %s", err.Error())
		return
	}

	events, err := bot.ParseRequest(c.Request)
	if err != nil {
		if err == linebot.ErrInvalidSignature {
			log.Print(err)
		}
		return
	}
	log.Printf("event: %+v", events)

	for _, event := range events {
		userID, roomID, groupID, userName := "", "", "", ""
		userID = event.Source.UserID
		roomID = event.Source.RoomID
		groupID = event.Source.GroupID

		// line develop webhook verify
		if userID == "Udeadbeefdeadbeefdeadbeefdeadbeef" {
			return
		}

		profile, err := bot.GetProfile(userID).Do()
		if err != nil {
			log.Fatal("GetProfile error %s", err.Error())
		} else {
			userName = profile.DisplayName
			log.Printf("userName = %s, StatusMessage = %s, PictureURL = %s", userName, profile.StatusMessage, profile.PictureURL)
		}

		log.Printf("template userID = %s, roomID = %s, GroupID = %s", userID, roomID, groupID)
		switch event.Type {
		case linebot.EventTypeMessage:
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				log.Printf("message.Text = %s", message.Text)

				switch message.Text {
				case "AdminMode":
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("你已進入為管理者模式！")).Do(); err != nil{
						log.Printf("change to Admin mode error: %+v", err)
					}

				case "GetLottery":
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("你獲得了一張彩票！")).Do(); err != nil{
						log.Printf("Reply error: %+v", err)
					}
				default:

				}
			case *linebot.ImageMessage:
				log.Printf("ImageMessage message.PreviewImageURL = %s, message.OriginalContentURL = %s", message.PreviewImageURL, message.OriginalContentURL)

			case *linebot.VideoMessage:
				log.Printf("VideoMessage message.PreviewImageURL = %s, message.OriginalContentURL = %s", message.PreviewImageURL, message.OriginalContentURL)

			case *linebot.AudioMessage:
				log.Printf("AudioMessage message.Duration = %d, message.OriginalContentURL = %s", message.Duration, message.OriginalContentURL)

			case *linebot.FileMessage:
				log.Printf("FileMessage message.FileName = %s, message.FileSize = %d", message.FileName, message.FileSize)

			case *linebot.LocationMessage:
				log.Printf("LocationMessage message.Latitude = %f, message.Longitude = %f", message.Latitude, message.Longitude)

			case *linebot.StickerMessage:
				log.Printf("StickerMessage message.PackageID = %s, message.StickerID = %s", message.PackageID, message.StickerID)

			default:
				log.Printf("Unknown message: %v", message)

			}

		case linebot.EventTypeJoin:
			log.Printf("Enter EventTypeJoin event = %s", event)

		case linebot.EventTypeFollow:
			log.Printf("Enter EventTypeFollow event = %s", event)

		case linebot.EventTypeUnfollow:
			log.Printf("Enter EventTypeUnfollow event = %s", event)

		case linebot.EventTypeLeave:
			log.Printf("Enter EventTypeLeave event = %s", event)

		case linebot.EventTypePostback:
			postMessage := event.Postback.Data

			//id#command#data
			if strings.Contains(postMessage, "#") {
				temp := strings.Split(postMessage, "#")
				if len(temp) >= 2 {
					switch temp[1] {
					case "postbackAction1":
						if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("觸發動作1")).Do(); err != nil{
							log.Printf("change to Admin mode error: %+v", err)
						}
					case "postbackAction2":
						if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("觸發動作2")).Do(); err != nil{
							log.Printf("change to Admin mode error: %+v", err)
						}
					}
				}
			}

		case linebot.EventTypeBeacon:
			log.Printf("Enter EventTypeBeacon event = %s", event)
		}
	}
}