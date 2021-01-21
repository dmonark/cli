package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var paymentCmd = &cobra.Command{
	Use:   "payments",
	Short: "list all the payments",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Printing all the payments")
		return nil
	},
}

var paymentCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "create payment",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Creating payments")
		return nil
	},
}

func init() {
	paymentCmd.AddCommand(paymentCreateCmd)
}
