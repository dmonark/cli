package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
)

var paymentId string

var paymentCmd = &cobra.Command{
	Use:   "payment",
	Short: "Payment Details for a payment id",
	PreRun: validateAuth,
	RunE: func(cmd *cobra.Command, args []string) error {
		response, error := ExecuteRequest("http://api.razorpay.in:28080/v1/payments/"+paymentId, http.MethodGet, nil)
		if error != nil {
			fmt.Println(error.Error())
		}
		var data map[string]interface{}

		fmt.Println("Printing the payment details for:- ", paymentId)

		json.Unmarshal(response, &data)

		fmt.Println(data)

		return nil
	},
}

var paymentListCmd = &cobra.Command{
	Use:   "payments",
	Short: "List of payments associated with this merchant",
	PreRun: validateAuth,
	RunE: func(cmd *cobra.Command, args []string) error {
		response, error := ExecuteRequest("http://api.razorpay.in:28080/v1/payments", http.MethodGet, nil)
		if error != nil {
			fmt.Println(error.Error())
		}
		var data map[string]interface{}

		fmt.Println("Listing the payments details")

		json.Unmarshal(response, &data)

		fmt.Println(data)

		return nil
	},
}

func init() {
	paymentCmd.Flags().StringVarP(&paymentId, "id", "i", "", "Payment Id")
	paymentCmd.MarkFlagRequired("id")
	
	paymentCmd.AddCommand(paymentListCmd)
}
