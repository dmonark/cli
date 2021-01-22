package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
)

type Invoice struct {
}

var InvoiceCmd = &cobra.Command{
	Use:    "invoice",
	Short:  "print invoice",
	PreRun: validateAuth,
	RunE: func(cmd *cobra.Command, args []string) error {
		printInvoice(args)
		return nil
	},
}
var InvoiceId string

func printInvoice(args []string) {
	var response []byte
	var err error
	if InvoiceId == "" {
		url := "https://api.razorpay.com/v1/invoices"
		response, err = ExecuteRequest(url, http.MethodGet, nil)
		if err != nil {
			fmt.Println(err)
		}
		var data map[string]interface{}
		json.Unmarshal(response, &data)
		structureInvoiceList(data)
	} else {
		url := "https://api.razorpay.com/v1/invoices/" + InvoiceId
		response, err = ExecuteRequest(url, http.MethodGet, nil)
		if err != nil {
			fmt.Println(err)
		}
		var data map[string]interface{}
		json.Unmarshal(response, &data)
		structureInvoiceById(data)
	}

	// result, _ := json.MarshalIndent(data, "", "  ")
	// output := pretty.Pretty(response)
	// result := pretty.Color(output, nil)
	// fmt.Println(string(result))
}
func init() {
	InvoiceCmd.Flags().StringVarP(&InvoiceId, "string", "i", "", "To Print invoice by id")
	rootCmd.AddCommand(InvoiceCmd)
}
