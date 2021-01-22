package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

var ppage int
var pid string

var paymentListCmd = &cobra.Command{
	Use:    "payment",
	Short:  "payment list",
	PreRun: validateAuth,
	RunE: func(cmd *cobra.Command, args []string) error {
		var items []interface{}
		if pid != "" {
			response, error := ExecuteRequest("http://0.0.0.0:28080/v1/payments/"+pid, http.MethodGet, nil)
			if error != nil {
				fmt.Println(error.Error())
				os.Exit(1)
			}

			var data map[string]interface{}
			json.Unmarshal(response, &data)

			items = append(items, data)
		} else {
			skip := (ppage - 1) * 10
			response, error := ExecuteRequest("http://0.0.0.0:28080/v1/payments?skip="+fmt.Sprintf("%v", skip), http.MethodGet, nil)
			if error != nil {
				fmt.Println(error.Error())
				os.Exit(1)
			}

			var data map[string]interface{}
			json.Unmarshal(response, &data)

			if data["count"].(float64) < 1 {
				fmt.Println("No Entity found")
				os.Exit(1)
			}

			items = data["items"].([]interface{})
		}

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"ID", "Order ID", "Amount", "Currency", "Method", "Status", "Receipt"})
		table.SetHeaderColor(
			tablewriter.Colors{tablewriter.Bold, tablewriter.FgGreenColor},
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
			tablewriter.Colors{tablewriter.Bold, tablewriter.FgCyanColor},
		)
		table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
		table.SetAlignment(tablewriter.ALIGN_LEFT)

		for _, v := range items {
			new_map := v.(map[string]interface{})
			row := []string{
				fmt.Sprintf("%v", new_map["id"]),
				fmt.Sprintf("%v", new_map["order_id"]),
				fmt.Sprintf("%v", new_map["amount"]),
				fmt.Sprintf("%v", new_map["currency"]),
				fmt.Sprintf("%v", new_map["method"]),
				fmt.Sprintf("%v", new_map["status"]),
				fmt.Sprintf("%v", new_map["receipt"]),
			}
			table.Append(row)
		}

		table.Render()

		return nil

	},
}

func init() {
	paymentListCmd.Flags().IntVarP(&ppage, "page", "p", 1, "Page number")
	paymentListCmd.Flags().StringVarP(&pid, "id", "i", "", "Payment ID")
}
