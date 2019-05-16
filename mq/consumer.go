package mq

import (
	"fmt"
	"github.com/streadway/amqp"
	"github.com/weeon/contract"
)

type Consumer struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	tag     string
	done    chan error

	debug bool

	fn     func([]byte)
	logger contract.Logger

	queueName string
}

func NewConsumer(amqpURI, queueName, ctag string, fn func([]byte), l contract.Logger) (*Consumer, error) {
	c := &Consumer{
		conn:      nil,
		channel:   nil,
		tag:       ctag,
		done:      make(chan error),
		logger:    l,
		fn:        fn,
		queueName: queueName,
	}

	var err error

	c.logger.Infof("dialing %q", amqpURI)
	c.conn, err = amqp.Dial(amqpURI)
	if err != nil {
		return nil, fmt.Errorf("Dial error : %s ", err)
	}

	go func() {
		fmt.Printf("closing: %s", <-c.conn.NotifyClose(make(chan *amqp.Error)))
	}()

	c.logger.Info("got Connection, getting Channel")
	c.channel, err = c.conn.Channel()
	if err != nil {
		return nil, fmt.Errorf("Channel: %s ", err)
	}

	c.Do()

	return c, nil
}

func (c *Consumer) Do() error {
	c.logger.Info("Queue bound to Exchange, starting Consume (consumer tag %q)", c.tag)
	deliveries, err := c.channel.Consume(
		c.queueName, // name
		c.tag,       // consumerTag,
		false,       // noAck
		false,       // exclusive
		false,       // noLocal
		false,       // noWait
		nil,         // arguments
	)
	if err != nil {
		return fmt.Errorf("Queue Consume fail : %s ", err)
	}

	go c.handle(deliveries, c.done)
	return nil
}

func (c *Consumer) handlerWatch() {
	for {
		select {
		case <-c.done:
			c.logger.Info("reconnect...")
			c.Do()
		}
	}
}

func (c *Consumer) handle(deliveries <-chan amqp.Delivery, done chan error) {
	for d := range deliveries {
		if c.debug {
			c.logger.Debugf(
				"got %dB delivery: [%v] %q",
				len(d.Body),
				d.DeliveryTag,
				d.Body,
			)
		}
		c.fn(d.Body)
		err := d.Ack(false)
		if err != nil {
			c.logger.Errorw("ask fail",
				"error", err,
			)
		}
	}
	c.logger.Info("handle: deliveries channel closed")
	done <- nil
}
