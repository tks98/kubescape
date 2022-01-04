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
		} else if args[0] == "scan-status" {
			err := imageScanStatusCmd.RunE(cmd, args)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(registryCmd)
	registryCmd.PersistentFlags().StringVarP(&registryInfo.Image, "image", "", "", "The image URL to be scanned")
	registryCmd.PersistentFlags().BoolVarP(&registryInfo.All, "all", "", false, "Scans all of the images in the image registry")
	registryCmd.PersistentFlags().StringVarP(&registryInfo.Name, "name", "", "", "Specify the name of the registry to authenticate with")
	registryCmd.PersistentFlags().StringVarP(&registryInfo.Username, "username", "u", "", "Username to log into the registry")
	registryCmd.PersistentFlags().StringVarP(&registryInfo.Password, "password", "p", "", "Password to log into the registry")
	registryCmd.PersistentFlags().StringVarP(&registryInfo.URL, "url", "", "", "URL of the registry to log into")
	registryCmd.PersistentFlags().StringVarP(&registryInfo.AuthType, "authType", "", "", "Type of authentication to use to login to the registry")
}
