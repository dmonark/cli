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
	Short: "Razorpay auth",
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

var merchantKey string
var merchantSecret string

var authLoginCmd = &cobra.Command{
	Use:   "login",
	Short: "login user with key and secret",
	Long:  "All Razorpay APIs are authorized using Basic Authorization. Basic authorization requires <YOUR_KEY_ID>, <YOUR_KEY_SECRET>",
	RunE: func(cmd *cobra.Command, args []string) error {
		os.Setenv("rzp_key", merchantKey)
		os.Setenv("rzp_secret", merchantSecret)
		syscall.Exec(os.Getenv("SHELL"), []string{os.Getenv("SHELL")}, syscall.Environ())
		fmt.Println("User Logged in")
		return nil
	},
}

var authLogoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "logout user",
	RunE: func(cmd *cobra.Command, args []string) error {
		os.Unsetenv("rzp_key")
		os.Unsetenv("rzp_secret")
		syscall.Exec(os.Getenv("SHELL"), []string{os.Getenv("SHELL")}, syscall.Environ())
		fmt.Println("User Logged out")
		return nil
	},
}

func init() {
	authLoginCmd.Flags().StringVarP(&merchantKey, "key", "k", "", "Razorpay merchant key")
	authLoginCmd.Flags().StringVarP(&merchantSecret, "secret", "s", "", "Razorpay merchant secret")
	authLoginCmd.MarkFlagRequired("key")
	authLoginCmd.MarkFlagRequired("secret")

	authCmd.AddCommand(authLoginCmd, authLogoutCmd)
}

func validateAuth(cmd *cobra.Command, args []string) {
	if os.Getenv("rzp_key") == "" || os.Getenv("rzp_secret") == "" {
		color.Red("auth failed!")
		color.Yellow("please run 'rzp auth login -k <YOUR_KEY> -s <YOUR_SECRET>'")
		os.Exit(1)
	}
}
