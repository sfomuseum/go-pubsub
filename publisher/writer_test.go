package publisher

import (
	"context"
	"io"
	"testing"
)

func TestPublisherWriteCloser(t *testing.T) {

	ctx := context.Background()

	p, err := NewPublisher(ctx, "stdout://")

	if err != nil {
		t.Fatalf("Failed to create new publisher, %v", err)
	}

	wr := NewWriter(p)

	mw := io.MultiWriter(wr)

	_, err = mw.Write([]byte("hello world\n"))

	if err != nil {
		t.Fatalf("Failed to write message, %v", err)
	}

	err = wr.Close()

	if err != nil {
		t.Fatalf("Failed to close writer, %v", err)
	}
}
