package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var changehostnameCmd = &cobra.Command{
	Use:   "change-hostname",
	Short: "Change a specified hostname",
	Long:  "Command accepts a new hostname and changes the previous one",

	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println("change-hostname command requires exactly one argument")
			return
		}

		newHostname := args[0]

		if err := changeHostname(newHostname); err != nil {
			fmt.Printf("Failed change hostname: %v", err)
			return
		}

		fmt.Printf("Successfully change hostname: %s", newHostname)
	},
}

func init() {
	rootCmd.AddCommand(changehostnameCmd)
}
