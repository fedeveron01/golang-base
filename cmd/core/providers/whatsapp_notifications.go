package providers

type WhatsappNotifications interface {
	SendNotification(phoneNumber string, message string) error
}
