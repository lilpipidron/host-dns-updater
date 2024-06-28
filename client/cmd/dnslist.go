/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// dnslistCmd represents the dnslist command
var dnslistCmd = &cobra.Command{
	Use:   "dnslist",
	Short: "show dns list",
	Long:  "show you all available DNS",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(getDNSList())
	},
}

func init() {
	rootCmd.AddCommand(dnslistCmd)
}
