package main

import (
	"fmt"
	"sync"
)

var (
	counter int
	mu      sync.Mutex
)

type Send interface {
	PushNotifications(msg string, wg *sync.WaitGroup)
}

type Email struct{}

func (e Email) PushNotifications(msg string, wg *sync.WaitGroup) {
	defer wg.Done() // Call Done when the goroutine finishes
	mu.Lock()
	counter++
	mu.Unlock()
	fmt.Println("Email message sent: ", msg)
}

type Text struct{}

func (t Text) PushNotifications(msg string, wg *sync.WaitGroup) {
	defer wg.Done() // Call Done when the goroutine finishes
	mu.Lock()
	counter++
	mu.Unlock()
	fmt.Println("Text message sent: ", msg)
}

type Teams struct{}

func (t Teams) PushNotifications(msg string, wg *sync.WaitGroup) {
	defer wg.Done() // Call Done when the goroutine finishes
	mu.Lock()
	counter++
	mu.Unlock()
	fmt.Println("Teams message sent: ", msg)
}

func SendNotifications(notifiers []Send, msg string) {
	var wg sync.WaitGroup
	wg.Add(len(notifiers)) // Add the number of notifiers to the WaitGroup

	for _, notifier := range notifiers {
		go notifier.PushNotifications(msg, &wg)
	}

	wg.Wait() // Wait for all goroutines to finish
}

func main() {
	notifiers := []Send{
		Teams{}, Text{}, Email{},
	}

	SendNotifications(notifiers, "testing")
	fmt.Println("Counter = ", counter)
}
