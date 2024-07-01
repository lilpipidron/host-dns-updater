package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "client",
	Short: "A client for interacting with the gRPC service",
	Long:  `A client for interacting with the gRPC service that allows changing hostname and managing DNS servers.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
