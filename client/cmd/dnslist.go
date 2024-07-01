package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var dnslistCmd = &cobra.Command{
	Use:   "dns-list",
	Short: "List of all DNS",
	Long:  "Command shows all available dns",

	Run: func(cmd *cobra.Command, args []string) {
		if list, err := getDNSList(); err != nil {
			fmt.Printf("fail to get DNS list: %v", err)
		} else {
			for _, dns := range list {
				fmt.Println(dns)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(dnslistCmd)
}
