package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/olekukonko/tablewriter"
)

func structureInvoiceList(data map[string]interface{}) {

	items := data["items"].([]interface{})

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Invoice no", "Customer ID", "Customer Name", "Order ID", "Amount", "Status", "Issued At", "Expired At"})
	table.SetHeaderColor(
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgGreenColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgGreenColor},
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
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgCyanColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgCyanColor},
	)
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetAlignment(tablewriter.ALIGN_LEFT)

	for _, v := range items {
		new_map := v.(map[string]interface{})
		customer := new_map["customer_details"].(map[string]interface{})
		row := []string{
			fmt.Sprintf("%v", new_map["id"]),
			fmt.Sprintf("%v", new_map["invoice_number"]),
			fmt.Sprintf("%v", new_map["customer_id"]),
			fmt.Sprintf("%v", customer["customer_name"]),
			fmt.Sprintf("%v", new_map["order_id"]),
			fmt.Sprintf("%v", new_map["amount"]),
			fmt.Sprintf("%v", new_map["status"]),
			fmt.Sprintf("%v", GetReadbleDate(new_map["issued_at"])),
			fmt.Sprintf("%v", GetReadbleDate(new_map["expired_at"])),
		}
		table.Append(row)
	}

	table.Render()
}

func GetReadbleDate(date interface{}) string {
	if date == nil {
		return ""
	}
	return time.Unix(int64(date.(float64)), 0).String()
}
