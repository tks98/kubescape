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

		if registryInfo.Image == "" && !registryInfo.All {
			return fmt.Errorf("the registry scan command requires a URL to the registry image (--image=<imageURL>) or the --all flag to scan all images")
		}

		// check that the registry login credentials are in the config
		localConfig := cautils.NewLocalConfig(getter.GetArmoAPIConnector(), scanInfo.Account)
		registryConfig := localConfig.GetRegistryConfig()
		if registryConfig.RegistryName == "" || registryConfig.RegistryURL == "" {
			return fmt.Errorf("no registry information found. You must run kubescape registry login to authenticate with the registry")
		}

		// handle a harbor registry scan request depending on if it was a scan all request, or a single image scan
		if registryConfig.RegistryName == cautils.Harbor.String() {
			harborRegistry, err := harbor.NewHarborRegistry(registryConfig.RegistryURL, registryConfig.Credentials)
			if err != nil {
				return err
			}

			if registryInfo.All {
				fmt.Println("Scanning all")
				result, err := harborRegistry.ScanAll()
				if err != nil {
					return err
				}

				fmt.Println(result)
				return nil
			}

			result, err := harborRegistry.ScanImage(registryInfo.Image)
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
