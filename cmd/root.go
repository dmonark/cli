package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "rzp",
	Short: "Razorpay is the only payments solution in India that allows businesses to accept, process and disburse payments with its product suite.",
}

func Execute() {
	rootCmd.AddCommand(paymentCmd, orderCmd, authCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
