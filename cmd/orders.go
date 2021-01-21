package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

var opage int
var olimit int

var orderListCmd = &cobra.Command{
	Use:    "order",
	Short:  "order list",
	PreRun: validateAuth,
	Args:   cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		var items []interface{}
		if len(args) == 1 {
			response, error := ExecuteRequest("http://0.0.0.0:28080/v1/orders/"+args[0], http.MethodGet, nil)
			if error != nil {
				fmt.Println(error.Error())
			}
			if response == nil {
				fmt.Println("Empty Response")
				os.Exit(1)
			}

			var data map[string]interface{}
			json.Unmarshal(response, &data)

			items = append(items, data)
		} else {
			skip := (opage - 1) * olimit
			response, error := ExecuteRequest("http://0.0.0.0:28080/v1/orders?skip="+fmt.Sprintf("%v", skip)+"&count="+fmt.Sprintf("%v", olimit), http.MethodGet, nil)
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

			items = data["items"].([]interface{})
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

		for _, v := range items {
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
	orderListCmd.Flags().IntVarP(&opage, "page", "p", 1, "Page number")
	orderListCmd.Flags().IntVarP(&olimit, "limit", "l", 10, "Number of result on one page")
}
