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
        if update.Message == nil {
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