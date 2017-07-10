package main

import "fmt"

func main() {
	msgCh := make(chan Message, 1)
	errCh := make(chan FailedMessage, 1)

	// define message to be sent
	// msg := Message{
	// 	To:      []string{"frodo@underhill.me"},
	// 	From:    "gandalf@whitecouncil.org",
	// 	Content: "SUCCESS Keep it secret, keep it safe.",
	// }

	// define a failure message
	// failedMessage := FailedMessage{
	// 	ErrorMessage:    "FAIL Message intercepted by black rider",
	// 	OriginalMessage: Message{},
	// }

	// put the error messages in their respective channels
	// msgCh <- msg
	// errCh <- failedMessage

	// listen on both channels simultaneously and react to whichever one has a message to work with
	// select block is overloaded to handle multiple programming tasks, like switching between channels
	select {
	case receivedMsg := <-msgCh:
		fmt.Println(receivedMsg.Content)
	case receivedError := <-errCh:
		fmt.Println(receivedError.ErrorMessage)
	default:
		// always evaluates to true, prevents deadlock
		fmt.Println("No messages received")
	}

}

// Message format
type Message struct {
	To      []string
	From    string
	Content string
}

type FailedMessage struct {
	ErrorMessage    string
	OriginalMessage Message
}
