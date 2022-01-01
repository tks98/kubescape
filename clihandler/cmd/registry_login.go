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

		if len(args) != 5 {
			return fmt.Errorf("the registry login command requires a name, URL, username, and password")
		}

		// update the local json config with the supplied registry information
		localConfig := cautils.NewLocalConfig(getter.GetArmoAPIConnector(), scanInfo.Account)
		credentials := make(map[string]string)
		credentials["username"] = args[3]
		credentials["password"] = args[4]
		err := localConfig.SetRegistryCredentials(credentials)
		if err != nil {
			return fmt.Errorf("problem setting registry credentials %s\n", err)
		}

		err = localConfig.SetRegistryURL(args[2])
		if err != nil {
			return fmt.Errorf("problem setting registry url %s\n", err)
		}

		err = localConfig.SetRegistryName(args[1])
		if err != nil {
			return fmt.Errorf("problem setting registry name %s\n", err)
		}

		return nil

	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
}
