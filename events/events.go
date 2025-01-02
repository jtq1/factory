package events

import "fmt"

type eventManager struct {
	ch          chan string
	subjectFunc map[string]func() error
}

func NewEventManager() *eventManager {
	return &eventManager{
		ch:          make(chan string),
		subjectFunc: make(map[string]func() error),
	}
}

func (em *eventManager) Subscribe(subject string, funct func() error) error {
	if _, ok := em.subjectFunc[subject]; !ok {
		em.subjectFunc[subject] = funct
		return nil
	}

	return fmt.Errorf("subject already exist. Subject: %v", subject)
}

func (em *eventManager) Run() {
	fmt.Printf("Event poller started")
	go func() {
		for subject := range em.ch {
			if _, ok := em.subjectFunc[subject]; !ok {
				fmt.Printf("error event subject doesn't exist. Subject %v\n", subject)
				continue
			}

			go func() {
				err := em.subjectFunc[subject]()
				if err != nil {
					fmt.Printf("error executing event %v. Error %v\n", subject, err)
				}
			}()
		}
	}()
}

func (em *eventManager) Send(subject string) {
	em.ch <- subject
}
