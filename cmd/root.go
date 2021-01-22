package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "rzp",
	Short: "Razorpay - Best Payment Gateway for Online Payments - India",
	Long:  "Razorpay is the only payments solution in India that allows businesses to accept, process and disburse payments with its product suite. It gives you access to all payment modes including credit card, debit card, netbanking, UPI and popular wallets including JioMoney, Mobikwik, Airtel Money, FreeCharge, Ola Money and PayZapp.",
}

func Execute() {
	rootCmd.AddCommand(paymentListCmd, orderListCmd, InvoiceCmd, authCmd, customerListCmd, disputeListCmd, dashboardListCmd, balanceCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
