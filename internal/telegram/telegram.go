package telegram

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"sushi-backend/config"
	"sushi-backend/internal/logger"
	"sushi-backend/utils"
)

type Telegram struct {
	logger logger.ILogger
	config config.IConfig
}

func NewTelegram(deps TelegramDependencies) *Telegram {
	return &Telegram{
		logger: deps.Logger,
		config: deps.Config,
	}
}

func (t *Telegram) SendMessageToChannel(ctx context.Context, chatId, message string, markdown bool) {
	t.logger.Debug("Sending message to telegram")

	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", t.config.TelegramBotToken())

	parseMode := "Markdown"

	if !markdown {
		parseMode = "HTML"
	}

	body, err := json.Marshal(map[string]string{
		"chat_id":    chatId,
		"text":       message,
		"parse_mode": parseMode,
	})

	utils.PanicIfError(err)

	req := utils.PanicIfErrorWithResultReturning(http.NewRequest(http.MethodPost, url, bytes.NewBuffer(body)))

	req = req.WithContext(ctx)

	req.Header.Set("Content-Type", "application/json")

	client := http.DefaultClient

	utils.PanicIfErrorIsNotContextErrorWithResult(client.Do(req))
}
