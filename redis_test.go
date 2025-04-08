package pubsub

import (
	"net/url"
	"testing"
)

func TestRedisConfigFromURL(t *testing.T) {

	tests := map[string][2]string{
		"redis://testing":                                   [2]string{"localhost:6379", "testing"},
		"redis://testing?host=127.0.0.1":                    [2]string{"127.0.0.1:6379", "testing"},
		"redis://testing?host=127.0.0.1&port=9736":          [2]string{"127.0.0.1:9736", "testing"},
		"redis://?host=127.0.0.1&port=9736&channel=example": [2]string{"127.0.0.1:9736", "example"},
	}

	for uri, expected := range tests {

		u, err := url.Parse(uri)

		if err != nil {
			t.Fatalf("Failed to parse URI '%s', %v", uri, err)
		}

		endpoint, channel, err := RedisConfigFromURL(u)

		if err != nil {
			t.Fatalf("Failed to derive config from URI '%s', %v", uri, err)
		}

		if endpoint != expected[0] {
			t.Fatalf("Unexpected endpoint for URI '%s'. Expected '%s' but got '%s'", uri, expected[0], endpoint)
		}

		if channel != expected[1] {
			t.Fatalf("Unexpected channel for URI '%s'. Expected '%s' but got '%s'", uri, expected[1], channel)
		}

	}
}
