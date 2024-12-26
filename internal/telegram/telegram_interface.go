package telegram

import "context"

type ITelegram interface {
	SendMessageToChannel(ctx context.Context, chatId, message string, markdown bool)
}
