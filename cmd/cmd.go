package cmd

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/pianzm/kafka-redpanda/cmd/consumer"
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "POD Notification Service",
		Short: "POD Notification - Backend Service",
		Long:  "POD Notification - API Gateway for POD Notification Service",
	}
)

func Execute() {
	godotenv.Load()
	rootCmd.AddCommand(consumer.ServeConsumerCmd())

	if err := rootCmd.Execute(); err != nil {
		os.Exit(-1)
	}

}
