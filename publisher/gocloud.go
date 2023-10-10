package publisher

import (
	"context"
	"fmt"
	"net/url"
	"sync"

	aa_session "github.com/aaronland/go-aws-session"
	"github.com/aws/aws-sdk-go/aws/session"
	"gocloud.dev/pubsub"
	"gocloud.dev/pubsub/awssnssqs"
)

type GoCloudPublisher struct {
	Publisher
	topic *pubsub.Topic
}

// In principle this could also be done with a sync.OnceFunc call but that will
// require that everyone uses Go 1.21 (whose package import changes broke everything)
// which is literally days old as I write this. So maybe a few releases after 1.21.
//
// Also, _not_ using a sync.OnceFunc means we can call RegisterSchemes multiple times
// if and when multiple gomail-sender instances register themselves.

var register_mu = new(sync.RWMutex)
var register_map = map[string]bool{}

func init() {

	ctx := context.Background()
	err := RegisterSchemes(ctx)

	if err != nil {
		panic(err)
	}
}

// RegisterGoCloudSchemes will explicitly register all the schemes associated with the `GoCloudPublisher` interface.
func RegisterGoCloudSchemes(ctx context.Context) error {

	register_mu.Lock()
	defer register_mu.Unlock()

	to_register := []string{
		"awssqs-creds",
	}

	for _, scheme := range pubsub.DefaultURLMux().TopicSchemes() {
		to_register = append(to_register, scheme)
	}

	for _, scheme := range to_register {

		_, exists := register_map[scheme]

		if exists {
			continue
		}

		err := RegisterPublisher(ctx, scheme, NewGoCloudPublisher)

		if err != nil {
			return fmt.Errorf("Failed to register blob writer for '%s', %w", scheme, err)
		}

		register_map[scheme] = true
	}

	return nil
}

func NewGoCloudPublisher(ctx context.Context, uri string) (Publisher, error) {

	u, err := url.Parse(uri)

	if err != nil {
		return nil, fmt.Errorf("Failed to parse URI, %w", err)
	}

	var topic *pubsub.Topic

	switch u.Scheme {
	case "awssqs-creds":

		q := u.Query()

		region := q.Get("region")
		credentials := q.Get("credentials")
		queue_url := q.Get("queue-url")

		cfg, err := aa_session.NewConfigWithCredentialsAndRegion(credentials, region)

		if err != nil {
			return nil, fmt.Errorf("Failed to create new session for credentials '%s', %w", credentials, err)
		}

		sess, err := session.NewSession(cfg)

		if err != nil {
			return nil, fmt.Errorf("Failed to create AWS session, %w", err)
		}

		// https://gocloud.dev/howto/pubsub/publish/#sqs-ctor

		topic = awssnssqs.OpenSQSTopic(ctx, sess, queue_url, nil)

	default:

		topic, err = pubsub.OpenTopic(ctx, uri)

		if err != nil {
			return nil, err
		}
	}

	pub := &GoCloudPublisher{
		topic: topic,
	}

	return pub, err
}

func (pub *GoCloudPublisher) Publish(ctx context.Context, str_msg string) error {

	msg := &pubsub.Message{
		Body: []byte(str_msg),
	}

	return pub.topic.Send(ctx, msg)
}

func (pub *GoCloudPublisher) Close() error {
	ctx := context.Background()
	return pub.topic.Shutdown(ctx)
}
