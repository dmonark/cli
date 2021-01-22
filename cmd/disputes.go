package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

var dpage int
var did string

var disputeListCmd = &cobra.Command{
	Use:    "dispute",
	Short:  "dispute list",
	PreRun: validateAuth,
	Args:   cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		var items []interface{}
		if pid != "" {
			response, error := ExecuteRequest("http://0.0.0.0:28080/v1/disputes/"+did, http.MethodGet, nil)
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
			skip := (dpage - 1) * 10
			response, error := ExecuteRequest("http://0.0.0.0:28080/v1/disputes?skip="+fmt.Sprintf("%v", skip), http.MethodGet, nil)
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
		table.SetHeader([]string{"ID", "Payment ID", "Amount", "Currency", "Method", "Status", "Reason Code"})
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
				fmt.Sprintf("%v", new_map["payment_id"]),
				fmt.Sprintf("%v", new_map["amount"]),
				fmt.Sprintf("%v", new_map["currency"]),
				fmt.Sprintf("%v", new_map["method"]),
				fmt.Sprintf("%v", new_map["status"]),
				fmt.Sprintf("%v", new_map["reason_code"]),
			}
			table.Append(row)
		}

		table.Render()

		return nil

	},
}

func init() {
	disputeListCmd.Flags().IntVarP(&dpage, "page", "p", 1, "Page number")
	disputeListCmd.Flags().StringVarP(&did, "id", "i", "", "Dispute ID")
}
