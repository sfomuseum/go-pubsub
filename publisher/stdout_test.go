package publisher

import (
	"context"
	"testing"
)

func TestSdtoutPublisher(t *testing.T) {

	uris := []string{
		"stdout://",
		"stdout://?newline=true",
		"stdout://?newline=false",
	}

	ctx := context.Background()

	for _, u := range uris {

		_, err := NewPublisher(ctx, u)

		if err != nil {
			t.Fatalf("Failed to parse '%s', %v", u, err)
		}
	}
}
