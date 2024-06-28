package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var adddnsCmd = &cobra.Command{
	Use:   "adddns",
	Short: "add dns",
	Long:  "The function accepts and adds dns, maybe several at once",
	Run: func(cmd *cobra.Command, args []string) {
		for _, dns := range args {
			if err := addDNS(dns); err != nil {
				log.Printf("Failed add dns %s: %s", dns, err.Error())
			} else {
				log.Printf("Successfully add dns %s", dns)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(adddnsCmd)
}
