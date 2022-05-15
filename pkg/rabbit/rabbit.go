package rabbit

import (
	"context"
	"fmt"

	"github.com/streadway/amqp"
)

const rabbitKey = "rabbit"

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

func get(ctx context.Context) ConnecitonProvider {
	v := ctx.Value(rabbitKey)
	c, ok := v.(ConnecitonProvider)
	if !ok {
		panic(fmt.Errorf(rabbitKey+" has wrong type: %[1]v %[1]T", v))
	}

	return c
}

func Get(ctx context.Context) ConnecitonProvider { return get(ctx) }
