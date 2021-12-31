package cmd

import (
	"fmt"
	"github.com/armosec/kubescape/cautils"
	"github.com/armosec/kubescape/cautils/getter"
	"github.com/armosec/kubescape/registryhandler/harbor"
	"github.com/spf13/cobra"
)

var registryScanCmd = &cobra.Command{
	Use:   "scan <imageURL>",
	Short: "Log into the image registry source",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {

		if len(args) != 2 {
			return fmt.Errorf("the registry scan command requires a URL to the registry image")
		}

		imageURL := args[1]

		// check that the registry login credentials are in the config
		localConfig := cautils.NewLocalConfig(getter.GetArmoAPIConnector(), scanInfo.Account)
		registryConfig := localConfig.GetRegistryConfig()
		if registryConfig.RegistryName == "" || registryConfig.RegistryURL == "" || registryConfig.Credentials == nil {
			return fmt.Errorf("no registry information found. You must run kubescape registry login to authenticate with the registry")
		}

		if registryConfig.RegistryName == "harbor" {
			harborRegistry, err := harbor.NewHarborRegistry(registryConfig.RegistryURL, registryConfig.Credentials["username"], registryConfig.Credentials["password"])
			if err != nil {
				return err
			}

			if imageURL == "all" {
				result, err := harborRegistry.ScanAll()
				if err != nil {
					return err
				}

				fmt.Println(result)
			}

			result, err := harborRegistry.ScanImage(imageURL)
			if err != nil {
				return err
			}

			fmt.Println(result)
		}

		return nil

	},
}

func init() {
	rootCmd.AddCommand(registryScanCmd)
}
