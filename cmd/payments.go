package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
)

var PaymentId string

var PaymentCmd = &cobra.Command{
	Use:   "payments",
	Short: "list payment details",
	PreRun: validateAuth,
	RunE: func(cmd *cobra.Command, args []string) error {
		printPayments(args)
		return nil
	},
}

func printPayments(args []string) {

	var response []byte

	var err error

	if PaymentId == "" {
		url := "http://api.razorpay.in:28080/v1/payments"
		response, err = ExecuteRequest(url, http.MethodGet, nil)
		fmt.Println("Listing the payments")
	} else{
		url := "http://api.razorpay.in:28080/v1/payments/"+PaymentId
		response, err = ExecuteRequest(url, http.MethodGet, nil)
		fmt.Println("Printing the payment details for:- ", PaymentId)
	}
	
	if err != nil {
		fmt.Println(err.Error())
	}
	var data map[string]interface{}

	json.Unmarshal(response, &data)
	result, _ := json.MarshalIndent(data, "", "  ")
	fmt.Println(string(result))
}

func init() {
	PaymentCmd.Flags().StringVarP(&PaymentId, "id", "i", "", "Payment Id")
	rootCmd.AddCommand(PaymentCmd)
}
