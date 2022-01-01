package cmd

import (
	"fmt"
	"github.com/armosec/kubescape/cautils"
	"github.com/spf13/cobra"
	"os"
)

var registryInfo cautils.RegistryInfo
var registryCmd = &cobra.Command{
	Use:   "registry <command>",
	Short: "Interact with an image registry source",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println(fmt.Errorf("the registry command requires 1 or more arguments"))
			os.Exit(1)
		}

		// execute the specified registry command
		if args[0] == "login" {
			err := loginCmd.RunE(cmd, args)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		} else if args[0] == "scan" {
			err := registryScanCmd.RunE(cmd, args)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(registryCmd)
	registryCmd.PersistentFlags().StringVarP(&registryInfo.Image, "image", "i", "", "The image URL to be scanned")
	registryCmd.PersistentFlags().BoolVarP(&registryInfo.All, "all", "a", false, "Scans all of the images in the image registry")
}
