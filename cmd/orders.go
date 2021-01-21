package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var orderCmd = &cobra.Command{
	Use:   "orders",
	Short: "list all the orders",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Printing all the orders")
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
