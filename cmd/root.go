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
<<<<<<< HEAD
	rootCmd.AddCommand(paymentListCmd, orderListCmd, invoiceListCmd, authCmd, customerListCmd, disputeListCmd)
=======
	rootCmd.AddCommand(paymentListCmd, orderListCmd, InvoiceCmd, authCmd, customerListCmd)
>>>>>>> 07fa5d86eb76c03ee9b96474faae28475d743d18
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
