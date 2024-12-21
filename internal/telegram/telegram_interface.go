package telegram

type ITelegram interface {
	SendMessageToChannel(chatId, message string, markdown bool)
}
