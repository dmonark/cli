package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/zalando/go-keyring"
)

var orderCmd = &cobra.Command{
	Use:   "orders",
	Short: "list all the orders",
	RunE: func(cmd *cobra.Command, args []string) error {
		service := "my-app"
		user := "anon"
		secret, err := keyring.Get(service, user)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(secret)
		return nil
	},
}

var orderCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "create order",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Creating orders")
		return nil
	},
}

func init() {
	orderCmd.AddCommand(orderCreateCmd)
}
