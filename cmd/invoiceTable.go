package cmd

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/olekukonko/tablewriter"
)

func structureInvoiceById(data map[string]interface{}) {

	// Invoice info table
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Invoice Number", "Customer ID", "Date", "Invoice ID"})
	table.SetHeaderColor(
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
	)
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetAlignment(tablewriter.ALIGN_LEFT)

	row := []string{
		fmt.Sprintf("%v", data["invoice_number"]),
		fmt.Sprintf("%v", data["customer_id"]),
		fmt.Sprintf("%v", GetReadbleDate(data["date"])),
		fmt.Sprintf("%v", data["id"]),
	}
	table.Append(row)
	table.Render()

	//Address table

	table = tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"", "Shipping Address", "Billing Address"})
	table.SetHeaderColor(
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgGreenColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgGreenColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgGreenColor},
	)
	table.SetColumnColor(
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgRedColor},
		tablewriter.Colors{tablewriter.Normal, tablewriter.FgCyanColor},
		tablewriter.Colors{tablewriter.Normal, tablewriter.FgCyanColor},
	)
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetAlignment(tablewriter.ALIGN_LEFT)

	billing_address := data["customer_details"].(map[string]interface{})["billing_address"]
	shipping_address := data["customer_details"].(map[string]interface{})["shipping_address"]

	row = []string{
		fmt.Sprintf("%v", "Name"),
		fmt.Sprintf("%v", data["customer_details"].(map[string]interface{})["name"]),
		fmt.Sprintf("%v", data["customer_details"].(map[string]interface{})["name"]),
	}
	table.Append(row)
	var billingAddr, shippingAddr string
	if billing_address == nil {
		billingAddr = "N/A"
	} else {
		billingAddr = (billing_address.(map[string]interface{})["line1"]).(string) + (data["customer_details"].(map[string]interface{})["billing_address"].(map[string]interface{})["line2"].(string))
	}
	if shipping_address == nil {
		shippingAddr = "N/A"
	} else {
		shippingAddr = (shipping_address.(map[string]interface{})["line1"]).(string) + (data["customer_details"].(map[string]interface{})["shipping_address"].(map[string]interface{})["line2"].(string))
	}
	row = []string{
		fmt.Sprintf("%v", "Address"),
		fmt.Sprintf("%v", billingAddr),
		fmt.Sprintf("%v", shippingAddr),
	}
	table.Append(row)

	var shipping_city, billing_city string
	if billing_address == nil {
		billing_city = "N/A"
	} else {
		billing_city = billing_address.(map[string]interface{})["city"].(string)
	}
	if shipping_address == nil {
		shipping_city = "N/A"
	} else {
		shipping_city = shipping_address.(map[string]interface{})["city"].(string)
	}

	row = []string{
		fmt.Sprintf("%v", "City"),
		fmt.Sprintf("%v", billing_city),
		fmt.Sprintf("%v", shipping_city),
	}
	table.Append(row)

	var shipping_state, billing_state string
	if billing_address == nil {
		billing_state = "N/A"
	} else {
		billing_state = billing_address.(map[string]interface{})["state"].(string)
	}
	if shipping_address == nil {
		shipping_state = "N/A"
	} else {
		shipping_state = shipping_address.(map[string]interface{})["state"].(string)
	}

	row = []string{
		fmt.Sprintf("%v", "State"),
		fmt.Sprintf("%v", billing_state),
		fmt.Sprintf("%v", shipping_state),
	}
	table.Append(row)

	row = []string{
		fmt.Sprintf("%v", "Email"),
		fmt.Sprintf("%v", data["customer_details"].(map[string]interface{})["email"]),
		fmt.Sprintf("%v", data["customer_details"].(map[string]interface{})["email"]),
	}
	table.Append(row)

	row = []string{
		fmt.Sprintf("%v", "Contact"),
		fmt.Sprintf("%v", data["customer_details"].(map[string]interface{})["contact"]),
		fmt.Sprintf("%v", data["customer_details"].(map[string]interface{})["contact"]),
	}

	table.Append(row)
	table.Render()

	//Orders table

	table = tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Name", "Currency", "Quantity", " Net Amount"})
	table.SetFooter([]string{"", "", "", "Total", strconv.FormatFloat(data["amount"].(float64), 'f', 6, 64)})
	table.SetFooterColor(tablewriter.Colors{}, tablewriter.Colors{}, tablewriter.Colors{},
		tablewriter.Colors{tablewriter.Bold},
		tablewriter.Colors{tablewriter.FgHiRedColor})
	table.SetHeaderColor(
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgGreenColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgGreenColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgGreenColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgGreenColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgGreenColor},
	)
	table.SetColumnColor(
		tablewriter.Colors{tablewriter.Normal, tablewriter.FgCyanColor},
		tablewriter.Colors{tablewriter.Normal, tablewriter.FgCyanColor},
		tablewriter.Colors{tablewriter.Normal, tablewriter.FgCyanColor},
		tablewriter.Colors{tablewriter.Normal, tablewriter.FgCyanColor},
		tablewriter.Colors{tablewriter.Normal, tablewriter.FgCyanColor},
	)
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	orders := data["line_items"].([]interface{})

	for _, v := range orders {
		new_map := v.(map[string]interface{})
		row := []string{
			fmt.Sprintf("%v", new_map["id"]),
			fmt.Sprintf("%v", new_map["name"]),
			fmt.Sprintf("%v", new_map["currency"]),
			fmt.Sprintf("%v", new_map["quantity"]),
			fmt.Sprintf("%v", new_map["net_amount"]),
		}
		table.Append(row)
	}

	table.Render()
}

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
