package main

import (
    "os"
    tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
    "fmt"
    //"strings"
    "reflect"
    //"log"
    "strconv"
)

// export TELEGRAM_APITOKEN=...(number from telegram bot)

func logFile(update tgbotapi.Update, file *os.File) {    
    file.WriteString("Date = " + strconv.Itoa(update.Message.Date))
    file.WriteString("\nChat ID = " + strconv.FormatInt(update.Message.Chat.ID, 10))
    file.WriteString("\nMessage ID = " + strconv.Itoa(update.Message.MessageID))
    file.WriteString("\nUser Name = " + update.Message.From.UserName)
    file.WriteString("\nLast Name = " + update.Message.From.LastName)
    file.WriteString("\nLanguage Code = " + update.Message.From.LanguageCode)
    file.WriteString("\n Text = " + update.Message.Text)
    if update.Message.Document != nil {
        file.WriteString("\n Detected Document")
        file.WriteString("\n \t Document ID = " + update.Message.Document.FileID)
        file.WriteString("\n \t Document name = " + update.Message.Document.FileName)
        file.WriteString("\n \t Document size = " + strconv.Itoa(update.Message.Document.FileSize) + (" byte"))
    } else if update.Message.Photo != nil {
        file.WriteString("\n Detected Photo")
        file.WriteString("\n \t Photo ID = " + update.Message.Photo[3].FileID)
        file.WriteString("\n \t Width = " + strconv.Itoa(update.Message.Photo[3].Width))
        file.WriteString("\n \t Height = " + strconv.Itoa(update.Message.Photo[3].Height))
        file.WriteString("\n \t FileSize = " + strconv.Itoa(update.Message.Photo[3].FileSize) + (" byte"))
    }

    file.WriteString("\n-------------------\n")
    fmt.Println("Log finished...")
}

func main() {
    bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_APITOKEN"))
    if err != nil {
        panic(err)
    }
    fmt.Println(reflect.TypeOf(bot))
    bot.Debug = true

    // Создайте новую структуру updateConfig со смещением, равным 0. Смещения используются 
    // для того, чтобы убедиться, что Telegram знает, что мы обработали предыдущие значения, 
    // и нам не нужно их повторять.
    updateConfig := tgbotapi.NewUpdate(0)

    // Сообщите Telegram, что мы должны ждать до 30 секунд при каждом запросе обновления. 
    // Таким образом, мы можем получать информацию так же быстро, как и при выполнении 
    // множества частых запросов, без необходимости отправлять почти столько же запросов.
    updateConfig.Timeout = 30

    // Начните опрос Telegram на предмет обновлений.
    updates := bot.GetUpdatesChan(updateConfig)

    //configuration := bot.FileConfig
    // Обработка обновлений
    for update := range updates {
        fmt.Println("Start updates.")
        fmt.Println(reflect.TypeOf(update))

        file, err := os.OpenFile("logTg" + strconv.FormatInt(update.Message.Chat.ID, 10) + ".txt", os.O_APPEND|os.O_WRONLY, 0600)
        if err != nil {
            newFile, err := os.Create("logTg" + strconv.FormatInt(update.Message.Chat.ID, 10) + ".txt")
            if err != nil {
                fmt.Println("Unable to create file logging work TG:", err) 
                os.Exit(1) 
            }
            logFile(update, newFile)
            defer newFile.Close() 
            fmt.Println("! File closed !") 
        }
        logFile(update, file)
        defer file.Close() 
        fmt.Println("! File closed !")

        // Telegram может отправлять множество типов обновлений в зависимости от того, 
        // чем занимается ваш бот. Пока мы хотим просмотреть только сообщения, чтобы 
        // отменить любые другие обновления.
        if update.Message.Document != nil {
            msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
            msg.Text = "Документ загружен"
            bot.Send(msg)

            fileConf := tgbotapi.FileConfig {
                FileID: update.Message.Document.FileID,
            }
            _, err := bot.GetFile(fileConf)
            //photo, err := bot.GetFile(fileConf)
	        if err != nil {
		        msg.Text = "Не удалось скачать документ"
                bot.Send(msg)
	        }
            links, err := bot.GetFileDirectURL(update.Message.Document.FileID)
            fmt.Println("-------" + links)

            EchoDocumentSent(bot, update)
        } else if update.Message.Photo != nil {
            msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
            msg.Text = "Фотография загружена"
            bot.Send(msg)
            
            fileConf := tgbotapi.FileConfig {
                FileID: update.Message.Photo[3].FileID,
            }
            _, err := bot.GetFile(fileConf)
            //photo, err := bot.GetFile(fileConf)
	        if err != nil {
		        msg.Text = "Не удалось скачать фотографию"
                bot.Send(msg)
	        }
            links, err := bot.GetFileDirectURL(update.Message.Photo[3].FileID)
            fmt.Println("-------" + links)

            EchoPhotoSent(bot, update)
        } else if update.Message == nil {
            continue
        } else {
            if update.Message.IsCommand() {
                msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
                switch update.Message.Command() {
                case "command1":
                    msg.Text = "Лучший телеграм бот в мире !"
                default:
                    msg.Text = "I don't know that command"
                }
                bot.Send(msg)
            }
            EchoMessageSent(bot, update)
        }
        fmt.Println("//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//-//")
    }
}

// sudo go mod init echo.go 
// sudo go mod tidy
// sudo go build -o test