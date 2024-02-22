package subscriber

import (
	"context"
	"fmt"
	"net/url"
	"strconv"

	"github.com/go-redis/redis/v8"
	"github.com/sfomuseum/go-pubsub"
)

type RedisSubscriber struct {
	Subscriber
	redis_client  *redis.Client
	redis_channel string
}

func init() {
	ctx := context.Background()
	RegisterRedisSubscribers(ctx)
}

func RegisterRedisSubscribers(ctx context.Context) error {
	return RegisterSubscriber(ctx, "redis", NewRedisSubscriber)
}

func NewRedisSubscriber(ctx context.Context, uri string) (Subscriber, error) {

	u, err := url.Parse(uri)

	if err != nil {
		return nil, fmt.Errorf("Failed to parse URI, %w", err)
	}

	q := u.Query()

	host := pubsub.REDIS_DEFAULT_HOST
	port := pubsub.REDIS_DEFAULT_PORT

	if q.Has("host") {
		host = q.Get("host")
	}

	if q.Has("port") {
		str_port := q.Get("port")

		v, err := strconv.Atoi(str_port)

		if err != nil {
			return nil, fmt.Errorf("Failed to parse ?port= parameter, %w", err)
		}

		port = v
	}

	channel := q.Get("channel")

	if channel == "" {
		return nil, fmt.Errorf("Empty or missing ?channel= parameter")
	}

	addr := fmt.Sprintf("%s:%d", host, port)

	redis_client := redis.NewClient(&redis.Options{
		Addr: addr,
	})

	s := &RedisSubscriber{
		redis_client:  redis_client,
		redis_channel: channel,
	}

	return s, nil
}

func (s *RedisSubscriber) Listen(ctx context.Context, messages_ch chan string) error {

	pubsub_client := s.redis_client.PSubscribe(ctx, s.redis_channel)
	defer pubsub_client.Close()

	for {

		i, err := pubsub_client.Receive(ctx)

		if err != nil {
			return fmt.Errorf("Failed to receive message, %w", err)
		}

		if msg, _ := i.(*redis.Message); msg != nil {
			messages_ch <- msg.Payload
		}
	}

	return nil
}

func (s *RedisSubscriber) Close() error {
	return s.redis_client.Close()
}
