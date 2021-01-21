package cmd

import (
	"fmt"
	"os"
	"syscall"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
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
		fmt.Println(os.Setenv("rzp_key", merchantKey))
		fmt.Println(os.Setenv("rzp_secret", merchantSecret))
		syscall.Exec(os.Getenv("SHELL"), []string{os.Getenv("SHELL")}, syscall.Environ())
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

func validateAuth() {
	if os.Getenv("rzp_key") == "" || os.Getenv("rzp_secret") == "" {
		color.Red("auth failed!")
		color.Yellow("please run 'rzp auth login -k <YOUR_KEY> -s <YOUR_SECRET>'")
		os.Exit(1)
	}
}
