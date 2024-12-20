package telegram

type ITelegram interface {
	SendMessageToOrdersChannel(message string)
}
