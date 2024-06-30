package cmd

import (
	"fmt"
	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
)

var dnslistCmd = &cobra.Command{
	Use:   "dnslist",
	Short: "show dns list",
	Long:  "show you all available DNS",

	Run: func(cmd *cobra.Command, args []string) {
		if list, err := getDNSList(); err != nil {
			log.Error("fail to get DNS list:", err)
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
