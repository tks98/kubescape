package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/armosec/kubescape/cautils"
	"github.com/armosec/kubescape/cautils/getter"
	"github.com/armosec/kubescape/registryhandler/harbor"
	"github.com/spf13/cobra"
)

var imageScanStatusCmd = &cobra.Command{
	Use:   "scan-status <imageURL>",
	Short: "Retrieve the vulnerability scan status of an image",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {

		if registryInfo.Image == "" {
			return fmt.Errorf("the registry scan-status command requires a URL to the registry image (--image=<imageURL>)")
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

			scanStatus, err := harborRegistry.GetImagesScanStatus(registryInfo.Image)
			if err != nil {
				return err
			}

			jsonBytes, err := json.MarshalIndent(scanStatus, "", "\t")
			if err != nil {
				return err
			}

			fmt.Println(string(jsonBytes))

		} else {
			return fmt.Errorf("registry %s is not yet supported by kubescape", registryConfig.RegistryName)
		}

		return nil

	},
}

func init() {
	rootCmd.AddCommand(imageScanStatusCmd)
}

