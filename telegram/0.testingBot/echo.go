package main

import (
	"fmt"
	"os"
	"reflect"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	//"log"
	"errors"
	"io"
	"net/http"
	"os/exec"
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

		func() {
			file, err := os.OpenFile("./users/directory_"+strconv.FormatInt(update.Message.Chat.ID, 10)+"/!LogTg"+strconv.FormatInt(update.Message.Chat.ID, 10)+".txt", os.O_APPEND|os.O_WRONLY, 0600)
			if err != nil {
				newFile, err := os.Create("./users/directory_" + strconv.FormatInt(update.Message.Chat.ID, 10) + "/!LogTg" + strconv.FormatInt(update.Message.Chat.ID, 10) + ".txt")
				if err != nil {
					errMkdir := os.Mkdir("./users/directory_"+strconv.FormatInt(update.Message.Chat.ID, 10), 0755)
					if errMkdir != nil {
						fmt.Printf("Ошибка при создании директории: %v\n", errMkdir)
						os.Exit(1)
					}
					newFileinDir, errFileinDir := os.Create("./users/directory_" + strconv.FormatInt(update.Message.Chat.ID, 10) + "/!LogTg" + strconv.FormatInt(update.Message.Chat.ID, 10) + ".txt")
					if errFileinDir != nil {
						fmt.Printf("Ошибка при создании файла в директории: %v\n", errMkdir)
						os.Exit(1)
					}
					defer newFileinDir.Close()
					logFile(update, newFileinDir)
					fmt.Println("! File closed !")
				}
				//err := os.Mkdir("./users/directory_"+strconv.FormatInt(update.Message.Chat.ID, 10), os.ModePerm) // 0755 - права доступа к созданной папке
				defer newFile.Close()
				logFile(update, newFile)
				fmt.Println("! File closed !")
			}
			defer file.Close()
			logFile(update, file)
			fmt.Println("! File closed !")
		}()

		// Telegram может отправлять множество типов обновлений в зависимости от того,
		// чем занимается ваш бот. Пока мы хотим просмотреть только сообщения, чтобы
		// отменить любые другие обновления.
		if update.Message.Document != nil {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
			msg.Text = "Документ загружен"
			bot.Send(msg)

			//photo, err := bot.GetFile(fileConf)
			links, errLink := bot.GetFileDirectURL(update.Message.Document.FileID)
			if errLink != nil {
				msg.Text = "Не удалось получить ссылку на Документ"
				bot.Send(msg)
			}
			fmt.Println("-------" + links)
			fileName := update.Message.Document.FileName
			errDow := downloadFile(links, "./users/directory_"+strconv.FormatInt(update.Message.Chat.ID, 10)+"/"+fileName)
			if errDow != nil {
				msg.Text = "Не удалось скачать документ"
				bot.Send(msg)
			}
			fmt.Printf("File %s downlaod in current working directory", fileName)

			EchoDocumentSent(bot, update)

			cmd := exec.Command("./re_formating.exe")
			cmdArgs := []string{"./users/directory_" + strconv.FormatInt(update.Message.Chat.ID, 10) + "/" + fileName}
			cmd.Args = append(cmd.Args, cmdArgs[0])
			cmd.Stdout = os.Stdout
			fmt.Println("Start check")
			err := cmd.Run()
			if err != nil {
				msg.Text = "Не удалось обработать файл " + fileName
				bot.Send(msg)
			} else {
				answerMsg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
				fileCheck, err := os.ReadFile("./users/directory_" + strconv.FormatInt(update.Message.Chat.ID, 10) + "/" + fileName + ".txt")
				if err != nil {
					fmt.Printf("Ошибка при открытии отладочной информации файла: %v\n", err)
					os.Exit(1)
				}
				answerMsg.Text = string(fileCheck)
				bot.Send(answerMsg)
			}
			fmt.Println("Finished check")
		} else if update.Message.Photo != nil {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
			msg.Text = "Фотография загружена"
			bot.Send(msg)

			//photo, err := bot.GetFile(fileConf)
			links, errLink := bot.GetFileDirectURL(update.Message.Photo[3].FileID)
			if errLink != nil {
				msg.Text = "Не удалось получить ссылку на Фото"
				bot.Send(msg)
			}
			fmt.Println("-------" + links)
			fileName := strconv.Itoa(update.Message.MessageID)
			errDow := downloadFile(links, "./users/directory_"+strconv.FormatInt(update.Message.Chat.ID, 10)+"/"+fileName)
			if errDow != nil {
				msg.Text = "Не удалось скачать фото"
				bot.Send(msg)
			}
			fmt.Printf("File %s downlaod in current working directory", fileName)

			EchoPhotoSent(bot, update)
		} else if update.Message.Audio != nil {
			audio := tgbotapi.NewMessage(update.Message.Chat.ID, "")
			audio.Text = "Обработка аудио недоступна"
			bot.Send(audio)
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

func downloadFile(URL, fileName string) error {
	response, err := http.Get(URL)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	if response.StatusCode != 200 {
		return errors.New("Received non 200 response code")
	}

	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}
	return nil
}

// sudo go mod init echo.go
// sudo go mod tidy
// sudo go build -o test
