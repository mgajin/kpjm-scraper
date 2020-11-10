package mq

import (
	"context"

	"time"

	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
)

// Client struct is wrapper around go-redis client.
// It is used as message-queue client.
type Client struct {
	*redis.Client
	Subscriptions map[string]*redis.PubSub
}

// Config struct holds client's configurations.
type Config struct {
	ConnectionURL string
	// DBUsername    string
	// DBPassword    string
}

// NewClient creates new Client.
// Returns *Client if error.
func NewClient(config *Config) (*Client, error) {

	options, err := redis.ParseURL(config.ConnectionURL)
	if err != nil {
		return nil, err
	}

	ctx := context.Background()
	client := redis.NewClient(options)
	if _, err = client.Ping(ctx).Result(); err != nil {
		return nil, err
	}

	subs := make(map[string]*redis.PubSub)
	redisClient := &Client{
		Client:        client,
		Subscriptions: subs,
	}

	return redisClient, nil
}

// Subscribe subscribes client to given channel name.
// Returns error.
func (c *Client) Subscribe(channelName string) error {

	if c.isSubscribed(channelName) {
		return errors.New("client already subscribed to channel")
	}

	sub := c.Client.Subscribe(context.Background(), channelName)
	message, err := sub.Receive(context.Background())

	if err != nil {
		return errors.Wrap(err, "failed to subscribe to channel "+channelName)
	} else if _, ok := message.(*redis.Subscription); !ok {
		return errors.New("subscription confirmation not received")
	}

	c.Subscriptions[channelName] = sub

	return nil
}

// Consume reads message from subscribed channel.
// Returns message content and error.
func (c *Client) Consume(channelName string) (string, error) {

	sub, exists := c.Subscriptions[channelName]
	if !exists {
		return "", errors.New("client isn't subscribed to channel " + channelName)
	}
	message := <-sub.Channel()

	return message.Payload, nil
}

// Publish sends new message to channel.
// Returns error.
func (c *Client) Publish(channelName, message string) error {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result := c.Client.Publish(ctx, channelName, message)

	return result.Err()
}

// isSubscribed checks if client is already subscribed to given channel.
func (c *Client) isSubscribed(channelName string) (exists bool) {
	_, exists = c.Subscriptions[channelName]
	return
}
