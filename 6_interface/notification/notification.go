package notification

// comment....
type Notifiable interface {
	Notify()
}

// comment..
func NotifyExternalFunc(n Notifiable) {
	n.Notify()
}
