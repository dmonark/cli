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

var cpage int
var climit int

var customerListCmd = &cobra.Command{
	Use:    "customer",
	Short:  "customer list",
	PreRun: validateAuth,
	Args:   cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		var items []interface{}
		if len(args) == 1 {
			response, error := ExecuteRequest("http://0.0.0.0:28080/v1/customers/"+args[0], http.MethodGet, nil)
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
			skip := (cpage - 1) * climit
			response, error := ExecuteRequest("http://0.0.0.0:28080/v1/customers?skip="+fmt.Sprintf("%v", skip)+"&count="+fmt.Sprintf("%v", climit), http.MethodGet, nil)
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
		table.SetHeader([]string{"ID", "Name", "Contact", "Email", "GSTIN", "Created At"})
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
				fmt.Sprintf("%v", new_map["name"]),
				fmt.Sprintf("%v", new_map["contact"]),
				fmt.Sprintf("%v", new_map["email"]),
				fmt.Sprintf("%v", new_map["gstin"]),
				fmt.Sprintf("%v", time.Unix(int64(new_map["created_at"].(float64)), 0)),
			}
			table.Append(row)
		}

		table.Render()

		return nil

	},
}

func init() {
	customerListCmd.Flags().IntVarP(&cpage, "page", "p", 1, "Page number")
	customerListCmd.Flags().IntVarP(&climit, "limit", "l", 10, "Number of result on one page")
}
