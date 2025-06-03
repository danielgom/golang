package main

import (
	"context"
	"fmt"
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-aws/sqs"
	"github.com/aws/aws-sdk-go-v2/config"
	amazonsqs "github.com/aws/aws-sdk-go-v2/service/sqs"
	transport "github.com/aws/smithy-go/endpoints"
	"github.com/samber/lo"
	"log"
	"log/slog"
	"net/url"
	"slices"
)

func main() {

	myIntList := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	myEvenList := make([]int, 0, len(myIntList))
	for _, in := range myIntList {
		if in%2 == 0 {
			myEvenList = append(myEvenList, in)
		}
	}

	slices.Clip(myEvenList)

	subscriber, err := New()
	if err != nil {
		log.Fatal(err)
	}

	if err = ProcessMessages(subscriber, "test-queue"); err != nil {
		log.Fatal(err)
	}
}

func New() (*sqs.Subscriber, error) {
	ctx := context.Background()
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return nil, err
	}

	logger := watermill.NewSlogLogger(slog.Default())
	subscriberConfig := sqs.SubscriberConfig{
		DoNotCreateQueueIfNotExists: false,
		AWSConfig:                   cfg,
	}

	subscriber, err := sqs.NewSubscriber(subscriberConfig, logger)
	if err != nil {
		return nil, fmt.Errorf("failed to create sqs subscriber: %w", err)
	}

	return subscriber, nil
}

func ProcessMessages(subscriber *sqs.Subscriber, queueName string) error {
	ctx := context.Background()
	messages, err := subscriber.Subscribe(ctx, queueName)
	if err != nil {
		return fmt.Errorf("failed to subscribe: %w", err)
	}

	for msg := range messages {
		slog.InfoContext(ctx, "received message", "body", string(msg.Payload))
		msg.Ack()
	}

	return nil
}

func sqsOptions() []func(options *amazonsqs.Options) {
	return []func(options *amazonsqs.Options){
		amazonsqs.WithEndpointResolverV2(sqs.OverrideEndpointResolver{
			Endpoint: transport.Endpoint{
				URI: *lo.Must(url.Parse("http://localhost:4566")),
			},
		}),
	}

}
