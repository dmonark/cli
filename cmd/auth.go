package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/zalando/go-keyring"
)

var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "shows who is login",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Printing login info")
		return nil
	},
}

var merchantKey string
var merchantSecret string

var authLoginCmd = &cobra.Command{
	Use:   "login",
	Short: "login user",
	RunE: func(cmd *cobra.Command, args []string) error {
		service := "my-app"
		user := "anon"
		password := "secret 12345"

		// set password
		err := keyring.Set(service, user, password)
		if err != nil {
			log.Fatal(err)
		}

		// get password
		secret, err := keyring.Get(service, user)
		if err != nil {
			log.Fatal(err)
		}

		log.Println(secret)

		fmt.Println("User Login")
		return nil
	},
}

var authLogoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "logout user",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("User Logout")
		return nil
	},
}

func init() {
	authLoginCmd.Flags().StringVarP(&merchantKey, "key", "k", "", "Key of merchant")
	authLoginCmd.Flags().StringVarP(&merchantSecret, "secret", "s", "", "Secret of merchant")
	authLoginCmd.MarkFlagRequired("key")
	authLoginCmd.MarkFlagRequired("secret")

	authCmd.AddCommand(authLoginCmd, authLogoutCmd)
}
