package main

import (
    //"os"
    tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
    "fmt"
    //"strings"
    //"reflect"
    //"log"
    //"strconv"
)

func EchoMessageSent(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
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
    fmt.Println("Message send!")
}