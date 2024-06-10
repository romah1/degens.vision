package mint_tracker_mints

import (
	"context"
)

const (
	kExchangeName         = "mint_tracker_mints_exchange"
	kQueueName            = "mint_tracker_mints"
	kExchangePublishDelay = 5000
)

type MessageType = string

type Message struct {
	ProgramId string
	Signature string
}

type MessageProcessor interface {
	Process(ctx context.Context, message *Message) error
}
