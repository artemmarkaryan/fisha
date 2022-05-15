package rabbit

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/artemmarkaryan/fisha-facade/pkg/logy"
	"github.com/streadway/amqp"
)

const rabbitKey = "rabbit"
const bufSize = 100

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
}

type ConnecitonProvider func() *amqp.Connection

func (c Config) dialString() string {
	return fmt.Sprintf("amqp://%v:%v@%v:%v/", c.User, c.Password, c.Host, c.Port)
}

func Init(ctx context.Context, cfg Config) (context.Context, error) {
	conn, err := amqp.Dial(cfg.dialString())
	if err != nil {
		return ctx, err
	}

	var cp ConnecitonProvider = func() *amqp.Connection { return conn }

	return context.WithValue(ctx, rabbitKey, cp), err
}

func Get(ctx context.Context) ConnecitonProvider {
	v := ctx.Value(rabbitKey)
	c, ok := v.(ConnecitonProvider)
	if !ok {
		panic(fmt.Errorf(rabbitKey+" has wrong type: %[1]v %[1]T", v))
	}

	return c
}

func withQ(ctx context.Context, qName string, f func(*amqp.Channel, amqp.Queue) error) error {
	conn := Get(ctx)()
	defer func() { _ = conn.Close() }()

	channel, err := conn.Channel()
	if err != nil {
		return err
	}

	defer func() { _ = channel.Close() }()

	q, err := channel.QueueDeclare(
		qName,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	return f(channel, q)
}

func Produce(ctx context.Context, qName string, data []byte) error {
	return withQ(ctx, qName, func(ch *amqp.Channel, q amqp.Queue) error {
		return ch.Publish(
			"",
			q.Name,
			false,
			false,
			amqp.Publishing{
				ContentType: "text/json",
				Body:        data,
			})
	})
}

func Consume[T any](ctx context.Context, qName string, stop chan struct{}) (chan T, error) {
	var objs = make(chan T, bufSize)

	err := withQ(ctx, qName, func(ch *amqp.Channel, q amqp.Queue) error {
		msgs, err := ch.Consume(
			q.Name,
			"",
			true,
			false,
			false,
			false,
			nil,
		)
		if err != nil {
			return err
		}

		go func() {
			select {
			case <-stop:
				return
			case d := <-msgs:
				var obj T
				if err = json.Unmarshal(d.Body, &obj); err != nil {
					logy.Log(ctx).Errorf("unknown msg format: %q", string(d.Body))
					break
				}

				objs <- obj
			}
		}()

		return nil
	})

	if err != nil {
		return nil, err
	}

	return objs, nil
}
