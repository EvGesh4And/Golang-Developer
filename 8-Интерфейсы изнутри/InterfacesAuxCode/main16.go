package main

type MsgUserBalanceChanged struct {
	userID  string
	balance string
}

type MsgEventChanged struct {
	eventID string
}

func processMessage(msg interface{}) {
	// Implement me.
}

/*
user "user-1" balance was changed to "1000"
event "event-1" was changed
unknown message: unknown
*/
func main() {
	processMessage(MsgUserBalanceChanged{"user-1", "1000"})
	processMessage(MsgEventChanged{"event-1"})
	processMessage("unknown")
}
