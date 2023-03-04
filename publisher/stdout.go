package publisher

import (
	"context"
	"os"
)

type StdoutPublisher struct {
	Publisher
}

func init() {
	ctx := context.Background()
	RegisterPublisher(ctx, "stdout", NewStdoutPublisher)
}

func NewStdoutPublisher(ctx context.Context, uri string) (Publisher, error) {

	pub := &StdoutPublisher{}

	return pub, nil
}

func (pub *StdoutPublisher) Publish(ctx context.Context, msg string) error {

	os.Stdout.Write([]byte(msg))
	return nil
}

func (pub *StdoutPublisher) Close() error {
	return nil
}
