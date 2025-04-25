package notifier

type Notifier interface {
	SendAlert(message string) error
}
