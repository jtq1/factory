package models

type EventService interface {
	Subscribe(string, func() error) error
	Send(string)
}
