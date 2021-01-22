package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

var cpage int
var cid string

var customerListCmd = &cobra.Command{
	Use:    "customer",
	Short:  "customer list",
	PreRun: validateAuth,
	RunE: func(cmd *cobra.Command, args []string) error {
		var items []interface{}
		if cid != "" {
			response, error := ExecuteRequest("http://0.0.0.0:28080/v1/customers/"+cid, http.MethodGet, nil)
			if error != nil {
				fmt.Println(error.Error())
				os.Exit(1)
			}

			var data map[string]interface{}
			json.Unmarshal(response, &data)

			items = append(items, data)
		} else {
			skip := (cpage - 1) * 10
			response, error := ExecuteRequest("http://0.0.0.0:28080/v1/customers?skip="+fmt.Sprintf("%v", skip), http.MethodGet, nil)
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
				fmt.Sprintf("%v", GetReadbleDate(new_map["created_at"])),
			}
			table.Append(row)
		}

		table.Render()

		return nil

	},
}

func init() {
	customerListCmd.Flags().IntVarP(&cpage, "page", "p", 1, "Page number")
	customerListCmd.Flags().StringVarP(&cid, "id", "i", "", "Customer ID")
}
