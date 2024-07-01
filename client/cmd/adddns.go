package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var adddnsCmd = &cobra.Command{
	Use:   "add-dns",
	Short: "Add a DNS to the list",
	Long:  "Command accepts one (or list of) DNS and adds them to the list",
	Run: func(cmd *cobra.Command, args []string) {
		for _, dns := range args {
			if err := addDNS(dns); err != nil {
				fmt.Printf("Failed add dns %s: %v", dns, err)
			} else {
				fmt.Printf("Successfully add dns %s", dns)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(adddnsCmd)
}
