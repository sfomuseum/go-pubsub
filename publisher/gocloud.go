package publisher

import (
	"context"
	"fmt"
	"net/url"

	aa_session "github.com/aaronland/go-aws-session"
	"github.com/aws/aws-sdk-go/aws/session"
	"gocloud.dev/pubsub"
	"gocloud.dev/pubsub/awssnssqs"
)

type GoCloudPublisher struct {
	Publisher
	topic *pubsub.Topic
}

func init() {

	ctx := context.Background()

	RegisterPublisher(ctx, "awsqs-creds", NewGoCloudPublisher)
	
	for _, scheme := range pubsub.DefaultURLMux().TopicSchemes() {

		err := RegisterPublisher(ctx, scheme, NewGoCloudPublisher)

		if err != nil {
			panic(err)
		}
	}
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
