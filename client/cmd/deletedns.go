package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var deletednsCmd = &cobra.Command{
	Use:   "deletedns",
	Short: "delete dns",
	Long:  "The function accepts DNS and removes it, can accept several",

	Run: func(cmd *cobra.Command, args []string) {
		for _, dns := range args {
			if err := deleteDNS(dns); err != nil {
				log.Printf("failed delete dns %s: %s", dns, err.Error())
			} else {
				log.Printf("successfully delte dns %s", dns)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(deletednsCmd)
}
