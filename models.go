package main

type Message struct {
	AccountID      string `json:"account_id"`
	MessageID      string `json:"message_id"`
	SenderNumber   string `json:"sender_number"`
	ReceiverNumber string `json:"receiver_number"`
}
