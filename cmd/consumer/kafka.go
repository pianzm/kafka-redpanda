package consumer

import (
	"context"
	"fmt"
	"time"

	"github.com/jaswdr/faker"
	"github.com/spf13/cobra"
	"github.com/twmb/franz-go/pkg/kgo"
)

var (
	KafkaCmd = &cobra.Command{
		Use:  "run-consumer",
		RunE: run,
	}
)

func ServeConsumerCmd() *cobra.Command {
	KafkaCmd.Flags().StringP("provider", "p", "", "provider, kafka or redpanda")
	return KafkaCmd
}

func run(cmd *cobra.Command, args []string) error {
	provider, _ := cmd.Flags().GetString("provider")
	var serverURL string
	if provider == "" {
		provider = "kafka"
	}
	// Create a new client
	if provider == "kafka" {
		serverURL = "localhost:9092"
	} else {
		serverURL = "127.0.0.1:61783"
	}

	seeds := []string{serverURL}

	cl, err := kgo.NewClient(
		kgo.SeedBrokers(seeds...),
		kgo.ConsumerGroup("my-group"),
		kgo.ConsumeTopics("bench"),
	)
	if err != nil {
		panic(err)
	}
	defer cl.Close()
	ctx := context.Background()

	// seed random data
	rec := faker.New()
	startProduces := time.Now()

	for i := 0; i < 1000000; i++ {
		record := &kgo.Record{Topic: "bench", Value: []byte(rec.Person().Name())}
		cl.Produce(ctx, record, func(_ *kgo.Record, err error) {
			if err != nil {
				fmt.Printf("record had a produce error: %v\n", err)
			}
		})
	}
	elapsesProducer := time.Since(startProduces)
	fmt.Printf("Produce %d records in %s\n", 1000000, elapsesProducer)

	var numOfMessage int64
	start := time.Now()
	// 2.) Consuming messages from a topic
	for numOfMessage < 1000000 {
		fetches := cl.PollFetches(ctx)
		if errs := fetches.Errors(); len(errs) > 0 {
			panic(fmt.Sprint(errs))
		}

		fetches.EachPartition(func(p kgo.FetchTopicPartition) {
			for range p.Records {
				numOfMessage++
			}
		})
	}
	elapse := time.Since(start)
	fmt.Printf("Consume %d records in %s\n", numOfMessage, elapse)

	return nil
}
