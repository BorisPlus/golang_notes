package main

import "fmt"

type MsgUserBalanceChanged struct {
	userID  string
	balance string
}

type MsgEventChanged struct {
	eventID string
}

func processMessage(msg interface{}) {

	switch message := msg.(type) {
	case MsgUserBalanceChanged:
		fmt.Printf("user %q balance was changed to %q\n", message.userID, message.balance)
	case MsgEventChanged:
		fmt.Printf("event %q was changed\n", message.eventID)
	default:
		fmt.Printf("unknown message: %q\n", message)
	}
}

func _dynamic() {
	processMessage(MsgUserBalanceChanged{"user-1", "1000"})
	processMessage(MsgEventChanged{"event-1"})
	processMessage("unknown")
}

// user "user-1" balance was changed to "1000"
// event "event-1" was changed
// unknown message: unknown
