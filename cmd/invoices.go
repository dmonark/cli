package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

var ipage int
var iid string

var invoiceListCmd = &cobra.Command{
	Use:    "invoice",
	Short:  "invoice list",
	PreRun: validateAuth,
	Args:   cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		var items []interface{}
		if iid != "" {
			response, error := ExecuteRequest("http://0.0.0.0:28080/v1/invoices/"+iid, http.MethodGet, nil)
			if error != nil {
				fmt.Println(error.Error())
				os.Exit(1)
			}
			if response == nil {
				fmt.Println("Empty Response")
				os.Exit(1)
			}

			var data map[string]interface{}
			json.Unmarshal(response, &data)

			items = append(items, data)
		} else {
			skip := (cpage - 1) * 10
			response, error := ExecuteRequest("http://0.0.0.0:28080/v1/invoices?skip="+fmt.Sprintf("%v", skip), http.MethodGet, nil)
			if error != nil {
				fmt.Println(error.Error())
				os.Exit(1)
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

			items = data["items"].([]interface{})
		}

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"ID", "Customer ID", "Order ID", "Amount", "Currency", "Created At"})
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

		for _, v := range items {
			new_map := v.(map[string]interface{})
			row := []string{
				fmt.Sprintf("%v", new_map["id"]),
				fmt.Sprintf("%v", new_map["customer_id"]),
				fmt.Sprintf("%v", new_map["order_id"]),
				fmt.Sprintf("%v", new_map["amount"]),
				fmt.Sprintf("%v", new_map["currency"]),
				fmt.Sprintf("%v", time.Unix(int64(new_map["created_at"].(float64)), 0)),
			}
			table.Append(row)
		}

		table.Render()

		return nil

	},
}

func init() {
	invoiceListCmd.Flags().IntVarP(&ipage, "page", "p", 1, "Page number")
	invoiceListCmd.Flags().StringVarP(&iid, "id", "i", "", "Invoice ID")
}
