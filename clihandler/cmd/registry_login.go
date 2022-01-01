package cmd

import (
	"fmt"
	"github.com/armosec/kubescape/cautils"
	"github.com/armosec/kubescape/cautils/getter"
	"github.com/spf13/cobra"
)

var loginCmd = &cobra.Command{
	Use:   "login <registryName> <registryUrl> <username> <password",
	Short: "Log into the image registry source",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {

		// update the local json config with the supplied registry information
		localConfig := cautils.NewLocalConfig(getter.GetArmoAPIConnector(), scanInfo.Account)
		credentials := make(map[string]string)
		credentials["username"] = registryInfo.Username
		credentials["password"] = registryInfo.Password
		err := localConfig.SetRegistryCredentials(credentials)
		if err != nil {
			return fmt.Errorf("problem setting registry credentials %s\n", err)
		}

		err = localConfig.SetRegistryURL(registryInfo.URL)
		if err != nil {
			return fmt.Errorf("problem setting registry url %s\n", err)
		}

		err = localConfig.SetRegistryName(registryInfo.Name)
		if err != nil {
			return fmt.Errorf("problem setting registry name %s\n", err)
		}
		fmt.Printf("Saved Harbor Login Credentials to %s", cautils.ConfigFileFullPath())

		return nil

	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
}
