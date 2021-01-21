package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
)

var orderId string

var orderCmd = &cobra.Command{
	Use:    "order",
	Short:  "fetch order by order Id",
	PreRun: validateAuth,
	RunE: func(cmd *cobra.Command, args []string) error {
		response, error := ExecuteRequest("http://api.razorpay.in:28080/v1/orders/"+orderId, http.MethodGet, nil)
		if error != nil {
			fmt.Println(error.Error())
		}
		var data map[string]interface{}

		fmt.Println(json.Unmarshal(response, &data))
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
	orderCmd.Flags().StringVarP(&orderId, "id", "i", "", "Order Id")
	orderCmd.MarkFlagRequired("id")
	orderCmd.AddCommand(orderCreateCmd)
}
