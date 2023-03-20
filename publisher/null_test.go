package publisher

import (
	"context"
	"testing"
)

func TestNullPublisher(t *testing.T) {

	uris := []string{
		"null://",
	}

	ctx := context.Background()

	for _, u := range uris {

		p, err := NewPublisher(ctx, u)

		if err != nil {
			t.Fatalf("Failed to parse '%s', %v", u, err)
		}

		err = p.Publish(ctx, "Test")

		if err != nil {
			t.Fatalf("Failed to publish message, %v", err)
		}
	}
}
