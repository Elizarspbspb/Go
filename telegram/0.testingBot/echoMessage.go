package main

import (
    "os"
    tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
    "fmt"
    //"strings"
    //"reflect"
    //"log"
    "strconv"
)

func Message1(bot *tgbotapi.BotAPI) {
    /*file, err := os.Create("logTgWoork.txt")
    if err != nil{
        fmt.Println("Unable to create file logging work TG:", err) 
        os.Exit(1) 
    }*/

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

    //file.WriteString(bot.Self.UserName)

    // Давайте рассмотрим каждое обновление, которое мы получаем от Telegram.
    for update := range updates {
        fmt.Println("Start Message 1")
        
        // Telegram может отправлять множество типов обновлений в зависимости от того, 
        // чем занимается ваш бот. Пока мы хотим просмотреть только сообщения, чтобы 
        // отменить любые другие обновления.
        if update.Message == nil {
            continue
        }

        //file, err := os.Create("logTgWoork.txt")
        //file, err := os.Create("logTg" + strconv.FormatInt(update.Message.Chat.ID, 10) + ".txt")
        file, err := os.OpenFile("logTg" + strconv.FormatInt(update.Message.Chat.ID, 10) + ".txt", os.O_APPEND|os.O_WRONLY, 0600)
        if err != nil{
            fmt.Println("Unable to create file logging work TG:", err) 
            os.Exit(1) 
        }
        file.WriteString("Chat ID = " + strconv.FormatInt(update.Message.Chat.ID, 10))
        file.WriteString("\nMessage ID = " + strconv.Itoa(update.Message.MessageID))
        file.WriteString("\nUser Name = " + update.Message.From.UserName)
        file.WriteString("\nText = " + update.Message.Text)
        file.WriteString("\n-------------------\n")

        // Теперь, когда мы знаем, что получили новое сообщение, мы можем составить ответ! 
        // Мы возьмем идентификатор чата и текст из входящего сообщения и используем его 
        // для создания нового сообщения.
            msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)

        // Мы также скажем, что это сообщение является ответом на предыдущее сообщение. 
        // Для любых других спецификаций, кроме идентификатора чата или текста, 
        // вам нужно будет задать поля в `MessageConfig`.
            msg.ReplyToMessageID = update.Message.MessageID

        // Хорошо, мы отправляем наше сообщение! Нас не волнует сообщение, 
        // которое мы только что отправили, поэтому мы его удалим.
        if _, err := bot.Send(msg); err != nil {
            // Обратите внимание, что паника - плохой способ обработки ошибок. 
            // В Telegram могут быть перебои в обслуживании или сетевые ошибки, вам следует 
            // повторить попытку отправки сообщений или более корректно обрабатывать сбои.
            panic(err)
        }
        fmt.Println("Check...")
        defer file.Close() 
        //file.WriteString("text")
    }
    //defer file.Close() 
    //file.WriteString("text")
    fmt.Println("Done.")
}