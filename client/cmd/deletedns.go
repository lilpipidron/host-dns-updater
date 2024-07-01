package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var deletednsCmd = &cobra.Command{
	Use:   "delete-dns",
	Short: "Deletes specified DNS",
	Long:  "Command accepts one (or list of) DNS and removes them from the list",

	Run: func(cmd *cobra.Command, args []string) {
		for _, dns := range args {
			if err := deleteDNS(dns); err != nil {
				fmt.Printf("failed delete dns %s: %v", dns, err)
			} else {
				fmt.Printf("successfully delete dns %s", dns)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(deletednsCmd)
}
