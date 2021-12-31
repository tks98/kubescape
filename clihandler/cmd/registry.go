package cmd

import (
	"fmt"
	"github.com/armosec/kubescape/cautils"
	"github.com/armosec/kubescape/cautils/getter"
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

		if args[0] == "login" {

			if len(args) != 4 {
				fmt.Println("the registry login command requires a URL, username, and password")
				os.Exit(1)
			}

			localConfig := cautils.NewLocalConfig(getter.GetArmoAPIConnector(), scanInfo.Account)
			credentials := make(map[string]string)
			credentials["username"] = args[2]
			credentials["password"] = args[3]
			err := localConfig.SetRegistryCredentials(credentials)
			if err != nil {
				fmt.Printf("problem setting registry credentials %s\n", err)
				os.Exit(1)
			}
			fmt.Println("Value added successfully.")
			err = localConfig.SetRegistryURL(args[1])
			if err != nil {
				fmt.Printf("problem setting registry url %s\n", err)
				os.Exit(1)
			}

		}

		localConfigPath := cautils.ConfigFileFullPath()
		fmt.Println(localConfigPath)

	},
}

func init() {
	rootCmd.AddCommand(registryCmd)
}
