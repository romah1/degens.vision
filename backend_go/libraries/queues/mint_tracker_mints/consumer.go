package mint_tracker_mints

import (
	"context"
	"encoding/json"

	amqp "github.com/rabbitmq/amqp091-go"
	"go.uber.org/zap"
)

type MintTrackerMintsConsumer struct {
	logger    *zap.SugaredLogger
	ch        *amqp.Channel
	messages  <-chan amqp.Delivery
	processor MessageProcessor
}

func NewMintTrackerMintsConsumer(
	logger *zap.SugaredLogger,
	ch *amqp.Channel,
	processor MessageProcessor,
) (*MintTrackerMintsConsumer, error) {
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
	err = ch.QueueBind(
		kQueueName,
		"user.event.publish",
		kExchangeName,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}
	q, err := ch.QueueDeclare(
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
	messages, err := ch.Consume(
		q.Name,
		"mint-tracker-mints-consumer",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}
	return &MintTrackerMintsConsumer{
		logger:    logger,
		ch:        ch,
		messages:  messages,
		processor: processor,
	}, nil
}

func (mc *MintTrackerMintsConsumer) Start(ctx context.Context) {
	mc.logger.Infof("Started processing queue; name=%s", kQueueName)
	for {
		select {
		case delivery, ok := <-mc.messages:
			if !ok {
				mc.logger.Errorf("failed to consume message; connection was closed")
				return
			}
			var message Message
			err := json.Unmarshal(delivery.Body, &message)
			if err != nil {
				mc.logger.Errorf("Failed to unmarhall delivery; raw=%s", string(delivery.Body))
				delivery.Ack(false)
				continue
			}
			mc.logger.Infof("Got message; program_id=%s; signature=%s", message.ProgramId, message.Signature)
			err = mc.processor.Process(ctx, &message)
			if err != nil {
				mc.logger.Errorf("Failed to process delivery; sig=%s; error=%s", message.Signature, err.Error())
				err = mc.ch.PublishWithContext(
					ctx,
					kExchangeName,
					"user.event.publish",
					false,
					false,
					amqp.Publishing{
						ContentType: "application/json",
						Body:        delivery.Body,
						Headers: amqp.Table{
							"x-delay": kExchangePublishDelay,
						},
					},
				)
				if err != nil {
					mc.logger.Errorf("failed to publish to exchange; err=%s", err.Error())
					err = delivery.Reject(true)
					if err != nil {
						mc.logger.Errorf("Failed to Reject delivery; error=%s", err.Error())
					}
				} else {
					err = delivery.Ack(false)
					if err != nil {
						mc.logger.Errorf("Failed to Ack delivery; error=%s", err.Error())
					} else {
						mc.logger.Infof("published to exchange; sig=%s", message.Signature)
					}
				}
			} else {
				err = delivery.Ack(false)
				if err != nil {
					mc.logger.Errorf("Failed to Ack delivery; error=%s", err.Error())
				}
			}
		case <-ctx.Done():
			return
		}
	}
}
