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
	QNames   []string
}

type Rabbit struct {
	*amqp.Channel
	q map[string]amqp.Queue
}

func (c Config) dialString() string {
	return fmt.Sprintf("amqp://%v:%v@%v:%v/", c.User, c.Password, c.Host, c.Port)
}

func Init(ctx context.Context, cfg Config) (context.Context, error) {
	conn, err := amqp.Dial(cfg.dialString())
	if err != nil {
		return ctx, fmt.Errorf("cant create connection: %w", err)
	}

	r := new(Rabbit)
	r.Channel, err = conn.Channel()
	if err != nil {
		return ctx, fmt.Errorf("cant get connection channel: %w", err)
	}

	r.q = make(map[string]amqp.Queue, len(cfg.QNames))
	for _, name := range cfg.QNames {
		r.q[name], err = r.Channel.QueueDeclare(
			name,
			false,
			false,
			false,
			false,
			nil,
		)
	}

	return context.WithValue(ctx, rabbitKey, *r), err
}

func Get(ctx context.Context) Rabbit {
	v := ctx.Value(rabbitKey)
	c, ok := v.(Rabbit)
	if !ok {
		panic(fmt.Errorf(rabbitKey+" has wrong type: %[1]v %[1]T", v))
	}

	return c
}

func Produce(ctx context.Context, qName string, data []byte) error {
	rabbit := Get(ctx)
	if _, ok := rabbit.q[qName]; !ok {
		return fmt.Errorf("queue named %q not declared", qName)
	}

	return rabbit.Publish(
		"",
		qName,
		true,
		false,
		amqp.Publishing{
			ContentType: "text/json",
			Body:        data,
		},
	)
}

func Consume[T any](ctx context.Context, qName string, stop chan struct{}) (chan T, error) {
	var objs = make(chan T, bufSize)

	rabbit := Get(ctx)
	if _, ok := rabbit.q[qName]; !ok {
		return nil, fmt.Errorf("queue named %q not declared", qName)
	}

	msgs, err := rabbit.Consume(
		qName,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("cant enable consumer: %w", err)
	}

	go func() {
		for {
			select {
			case <-stop:
				logy.Log(ctx).Infoln("shutting consumer...")
				return
			case d := <-msgs:
				var obj T
				if err = json.Unmarshal(d.Body, &obj); err != nil {
					logy.Log(ctx).Errorf("unknown msg format: %q", string(d.Body))
					break
				}

				objs <- obj
			}
		}
	}()

	return objs, nil
}
