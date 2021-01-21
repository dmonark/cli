package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/olekukonko/tablewriter"
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

var orderListCmd = &cobra.Command{
	Use:    "list",
	Short:  "order list",
	PreRun: validateAuth,
	RunE: func(cmd *cobra.Command, args []string) error {
		response, error := ExecuteRequest("http://0.0.0.0:28080/v1/orders", http.MethodGet, nil)
		if error != nil {
			fmt.Println(error.Error())
		}
		if response == nil {
			fmt.Println("Empty Response")
			os.Exit(1)
		}
		var data map[string]interface{}

		json.Unmarshal(response, &data)

		if data["count"].(float64) < 1 {
			fmt.Println("No Entity found")
			os.Exit(1)
		}

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"ID", "Amount", "Currency", "Status", "Receipt", "Amount paid"})
		table.SetHeaderColor(
			tablewriter.Colors{tablewriter.Bold, tablewriter.FgGreenColor},
			tablewriter.Colors{tablewriter.Bold, tablewriter.FgGreenColor},
			tablewriter.Colors{tablewriter.Bold, tablewriter.FgGreenColor},
			tablewriter.Colors{tablewriter.Bold, tablewriter.FgGreenColor},
			tablewriter.Colors{tablewriter.Bold, tablewriter.FgGreenColor},
			tablewriter.Colors{tablewriter.Bold, tablewriter.FgGreenColor},
		)
		table.SetColumnColor(
			tablewriter.Colors{tablewriter.Bold, tablewriter.FgCyanColor},
			tablewriter.Colors{tablewriter.Bold, tablewriter.FgCyanColor},
			tablewriter.Colors{tablewriter.Bold, tablewriter.FgCyanColor},
			tablewriter.Colors{tablewriter.Bold, tablewriter.FgCyanColor},
			tablewriter.Colors{tablewriter.Bold, tablewriter.FgCyanColor},
			tablewriter.Colors{tablewriter.Bold, tablewriter.FgCyanColor},
		)
		table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
		table.SetAlignment(tablewriter.ALIGN_LEFT)

		for _, v := range data["items"].([]interface{}) {
			new_map := v.(map[string]interface{})
			row := []string{
				fmt.Sprintf("%v", new_map["id"]),
				fmt.Sprintf("%v", new_map["amount"]),
				fmt.Sprintf("%v", new_map["currency"]),
				fmt.Sprintf("%v", new_map["status"]),
				fmt.Sprintf("%v", new_map["receipt"]),
				fmt.Sprintf("%v", new_map["amount_paid"]),
			}
			table.Append(row)
		}

		table.Render()

		return nil
	},
}

func init() {
	orderCmd.Flags().StringVarP(&orderId, "id", "i", "", "Order Id")
	orderCmd.MarkFlagRequired("id")
	orderCmd.AddCommand(orderListCmd)
}
