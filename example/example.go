package main

import (
	"context"
	// "errors"
	"fmt"

	"github.com/ralstan-vaz/observer"
)

type sendSMS struct {
}

func (c sendSMS) OnNotify(ctx context.Context, e observer.Event) error {
	fmt.Println(e.Name, e.Message.(string), "sendSMS")
	return nil
}

type sendEmail struct {
}

func (c sendEmail) OnNotify(ctx context.Context, e observer.Event) error {
	fmt.Println(e.Name, e.Message.(string), "sendEmail")
	return nil
}

type notconfirmed struct {
}

func (c notconfirmed) OnNotify(ctx context.Context, e observer.Event) error {
	fmt.Println(e.Name, e.Message.(string), "notconfirmed")
	return nil
}

func main() {
	observer.Observer.Register("confirmed", sendSMS{})
	observer.Observer.Register("confirmed", sendSMS{})
	observer.Observer.Register("confirmed", sendEmail{})
	observer.Observer.Register("confirmed", notconfirmed{})

	observer.Observer.Register("notconfirmed", notconfirmed{})
	ctx := context.Background()
	err := observer.Observer.Publish(ctx, "confirmed", "message")
	if err != nil {
		fmt.Println("Error", err)
	}
	observer.Observer.Publish(ctx, "notconfirmed", "45")
}
