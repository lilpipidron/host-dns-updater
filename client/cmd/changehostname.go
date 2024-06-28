package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var changehostnameCmd = &cobra.Command{
	Use:   "changehostname",
	Short: "The command change hostname",
	Long:  "The command accepts the new hostname and changes the previous one to the new one",

	Run: func(cmd *cobra.Command, args []string) {
		newHostname := args[0]

		if err := changeHostname(newHostname); err != nil {
			log.Fatalf("Failed change hostname: %s", err.Error())
		}

		fmt.Println("Successfully change hostname:", newHostname)
	},
}

func init() {
	rootCmd.AddCommand(changehostnameCmd)
}
