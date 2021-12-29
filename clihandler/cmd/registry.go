package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var registryCmd = &cobra.Command{
	Use:   "registry <command>",
	Short: "Interact with an image registry source",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println(fmt.Errorf("the registry command requires 1 or more arguments"))
			os.Exit(1)
		}

		fmt.Printf("doing registry stuff with args %v", args)
	},
}

func init() {
	rootCmd.AddCommand(registryCmd)
}
