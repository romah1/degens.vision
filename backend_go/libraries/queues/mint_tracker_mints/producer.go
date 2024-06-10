package mint_tracker_mints

import (
	"context"
	"encoding/json"

	amqp "github.com/rabbitmq/amqp091-go"
	"go.uber.org/zap"
)

type MintTrackerMintsQueueProducer struct {
	logger *zap.SugaredLogger
	ch     *amqp.Channel
}

func NewMintTrackerMintsProducer(logger *zap.SugaredLogger, ch *amqp.Channel) (*MintTrackerMintsQueueProducer, error) {
	err := ch.ExchangeDeclare(
		kExchangeName,
		"x-delayed-message",
		true,
		false,
		false,
		false,
		amqp.Table{
			"x-delayed-type": "direct",
		},
	)
	if err != nil {
		return nil, err
	}
	_, err = ch.QueueDeclare(
		kQueueName,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}
	return &MintTrackerMintsQueueProducer{
		logger: logger,
		ch:     ch,
	}, nil
}

func (mp *MintTrackerMintsQueueProducer) Produce(ctx context.Context, message *Message) error {
	body, err := json.Marshal(message)
	if err != nil {
		return err
	}
	err = mp.ch.PublishWithContext(
		ctx,
		"",
		kQueueName,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
	if err != nil {
		mp.logger.Errorf("Failed to publish message to queue; name=%s; error=%s", kQueueName, err.Error())
		return err
	}
	mp.logger.Infof("Published message to queue; name=%s", kQueueName)
	return nil
}
